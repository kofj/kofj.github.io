title: boot2docker共享目录unicode文件名支持问题修复
tags:
  - docker
  - boot2docker
categories:
  - 术业
  - docker
date: 2015-01-23 01:08:03
---

&emsp;&emsp;这是一篇补记,问题在两周之前就已经解决了,并且给官方发送了pull request.目前官方只是把这个PR添加到了v1.5.0版的milestone当中,如果你需要.可以看我fork修改后的代码,或者直接下载生成好的[ISO镜像](http://pan.baidu.com/s/1hqKiLgw).
### 0x1 问题
&emsp;&emsp;Docker是个非常不错的工具,但是源于其本身的实现原理(依赖于Linux的LXC)的原因,Windows和OS X系统的用户是没有办法使用原生的docker的.幸运的是,docker官方为我们这一类`非Linux用户`提供了一个很好的工具——boot2docker.<!-- more -->
&emsp;&emsp;正如官方项目描述所言,boot2docker实质就是一个轻量级的支持Docker的Linux系统,事实上它就是基于`Tiny Core Linux`这一五脏俱全的小麻雀定制而成的.boot2docker虽是好用,但是在实际运用当中我却遇到了麻烦:boot2docker中的容器无法正常处理文件名使用unicode编码的文件的问题(注意:是文件名编码为unicode,而不是文件内容).
&emsp;&emsp;或者简单点说,这个问题就是boot2docker中无法显示(`ls -al`)出自动挂载的VBox共享目录中的中文文件名的文件.
### 0x2 溯源
&emsp;&emsp;存在即合理,既然必须要在docker容器当中处理中文文件名的文件,那么这个问题就应该解决掉.找准病因,才能对症下药.面对困难我们应当大胆假设,小心求证.
&emsp;&emsp;首先怀疑是docker官方提供的image不支持中文,启动一个容器使用`locale`查看系统默认编码并不是GBK或者UTF-8.
![docker CentOS 默认编码](https://ww2.sinaimg.cn/large/6816152bgw1eojcme1s2rj20fa0a6dgp.jpg)
&emsp;&emsp;回想起来,当时这个假设还是相当的粗放,大胆的,根本没有仔细思考为什么是`不显示文件`而非`文件名乱码`.无知无畏,既然有了假设,不管这个假设正确与否,只有实践验证方知结果.通过`localedef`生成中文locale文件,并设置系统默认编码集为UTF-8.
![修改系统locale和编码集](https://ww2.sinaimg.cn/large/6816152bgw1eojce2nyzrj20fa0a6dh1.jpg)
&emsp;&emsp;但是,结果还是让人让人失望的,依然无法显示中文名的文件.
![试验失败](https://ww3.sinaimg.cn/large/6816152bgw1eojedcbf5kj20nu05wgm9.jpg)
&emsp;&emsp;既然在容器中修改系统编码,同样是无法列出中文名的文件.那只能猜测是boot2docker这个轻量级的Linux系统存在问题了.
![boot2docker中无法列出VBox共享目录中的中文文件](https://ww1.sinaimg.cn/large/6816152bgw1eojek6sk49j20nu0ci0uy.jpg)
&emsp;&emsp;通过ssh登录到boot2docker,在Home目录当中创建和查看中文名文件,验证了boot2docker本身是能够处理中文名文件的.而在VBox挂载进来的共享目录当中却是不能正常处理中文名文件.这说明问题是出在了boot2docker挂载共享目录这一环节.
![在boot2docker的home目录中测试中文名文件](https://ww4.sinaimg.cn/large/6816152bgw1eojez26e6yj20cf05it95.jpg)
### 0x3 除病
&emsp;&emsp;既然找到了病症所在,那我们就能够开始下药方了.通过阅读boot2docker代码得知,它是使用`rootfs/rootfs/etc/rc.d/automount-shares`这一脚本自动挂载共享目录的.脚本中核心函数如下:

```BASH
	# try mounting "$name" (which defaults to "$dir") at "$dir",
	# but quietly clean up empty directories if it fails
	try_mount_share() {
		dir="$1"
		name="${2:-$dir}"
		
		mkdir -p "$dir" 2>/dev/null
		if ! mount -t vboxsf -o "$mountOptions" "$name" "$dir" 2>/dev/null; then
			rmdir "$dir" 2>/dev/null || true
			while [ "$(dirname "$dir")" != "$dir" ]; do
				dir="$(dirname "$dir")"
				rmdir "$dir" 2>/dev/null || break
			done
			
			return 1
		fi
		
		return 0
	}
```

&emsp;&emsp;boot2docker是基于`Tiny Core Linux`构建的,而Tiny Core Linux的rootfs是使用`busybox`构建的.并且Tiny Core Linux中mount工具源自busybox工具集,挂载共享目录时的默认编码并非UTF-8,这就导致了在宿主机中使用UTF-8编码的中文名文件挂载进boot2docker后无法显示出来.
&emsp;&emsp;幸运的是,boot2docker集成了VBox的Gust Additions,其中包含了共享目录的专用挂载工具`mount.vboxsf`.我们可以通过`mount.vboxsf`的参数`iocharset`来设定挂载共享目录的编码集为UTF-8.修改后的shell脚本代码如下:
```BASH
	# try mounting "$name" (which defaults to "$dir") at "$dir",
	# but quietly clean up empty directories if it fails
	try_mount_share() {
		dir="$1"
		name="${2:-$dir}"
		
		mkdir -p "$dir" 2>/dev/null
		if ! mount.vboxsf -o "$mountOptions" "$name" "$dir" 2>/dev/null; then
			rmdir "$dir" 2>/dev/null || true
			while [ "$(dirname "$dir")" != "$dir" ]; do
				dir="$(dirname "$dir")"
				rmdir "$dir" 2>/dev/null || break
			done
			
			return 1
		fi
		
		return 0
	}
```
&emsp;&emsp;最后,让我们一起来看一下是使用修改了`automount-shares`的boot2docker中对中文名文件的支持情况:
![能够显示中文名文件,虽然是乱码](https://ww4.sinaimg.cn/large/6816152bgw1eojium0q9qj20il0ci0uy.jpg)
&emsp;&emsp;目前已经能够显示中文名文件,虽然是乱码,但是并不影响容器当中处理文件.
### 0x4 结语
&emsp;&emsp;通过boot2docker的Dockfile来构建镜像是比较耗费资源的,而且由于大陆网络的原因,还有可能无法下载构建过程中的某些包.为了方便大家使用,我在本地重新构建了一个修改自动挂载脚本后的boot2docker的ISO镜像上传到了百度网盘.
下载地址:[共享目录支持中文文件名的boot2dokcer镜像](http://pan.baidu.com/s/1hqKiLgw)
密码: 5qoc