title: 使用git archive增量包
date: 2015-03-16 16:22:26
tags:
- git
- 归档
- 增量
categories: 术业
---
&emsp;&emsp;在没有项目版本仓库的使用权限的情况时,或者客户的生产环境没有自动化部署工具时,我们在修改了程序后需要生成一个增量包给同事/客户.逐个的检出文件到压缩包里面是个笨办法,若你马虎一下,便可能会落下个文件.保持代码的目录结构也是个麻烦事儿.使用`git archive`生成增量包则要方便很多.当然,前提是你本地有用git管理代码版本.
<!-- more -->
&emsp;&emsp;如果你的代码没有使用git做版本管理,修改代码之前,可以代码目录创建一个本地仓库.
```bash
#cd SourcePath #进入项目目录
git init #在当前目录创建一个新的空的本地仓库
git add . #把当前目录下的所有文件全部添加到暂存区
git commit -m 'Project init.' #提交创建
```
&emsp;&emsp;此后,就能方便的使用`git archive`生成增量包了.使用示例如下,`dbce55b`是代码修改前的版本号,`59b2e92`是代码修改完成后的版本号:
```bash
git archive -o update.zip HEAD $(git diff dbce55b...59b2e92 --name-only)
```

