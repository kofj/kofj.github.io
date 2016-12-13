title: 永久删除git仓库中的文件与历史记录
date: 2015-03-02 15:14:15
tags:
  - git
categories: 术业
---
&emsp;&emsp;之前托管在git@osc的私有项目`神马磁力搜索`想要开源,作为个人技能展示作品.这就遇到了一个问题,脱离敏感数据.通过`git rm file`肯定是不能够满足这一需求的,git历史记录中还是能够查看被删除的内容的,导致敏感数据的泄露.那么,就必须要彻底删除git的历史记录了.
<!-- more -->
## 0x1 从仓库总清除文件
&emsp;&emsp;打开bash,进入项目目录,执行以下命令:
```shell
git filter-branch --force --index-filter 'git rm --cached --ignore-unmatch path-of-remove-file' --prune-empty --tag-name-filter cat -- --all
```
&emsp;&emsp;`path-of-remove-file`是你要删除的文件的相对目录(相对于git repo的根目录),可以使用通配符`*`匹配文件进行批量删除.当看到类似于下面的提示信息说明删除成功了:
```shell
Rewrite 59b2e9e1bdc898daa52085648d3a8def767560dc (429/429)
# Ref 'refs/heads/master' was rewritten
```
&emsp;&emsp;如果执行上述命令后出现的提示信息中有`unchanged`字样,说明repo中没有找到`path-of-remove-file`,请仔细检查路径和文件名是否正确.
## 0x2 推送修改结果
```shell
git push origin master --force #强制覆盖
```
## 0x3 回收磁盘空间
&emsp;&emsp;经过上述操作后,我们已经删除了文件,但是本地仓库中任然保留着这些objects.我们需要使用GC命令进行垃圾回收,彻底清除这些文件,回收磁盘空间.`#`后面为执行结果,我们可以看见.git目录大小明显缩小了.
```shell
rm -rf .git/refs/original/
du -sh .git
# 11M	.git
git reflog expire --expire=now --all
git gc --prune=now
#Counting objects: 3370, done.
#Delta compression using up to 4 threads.
#Compressing objects: 100% (2047/2047), done.
#Writing objects: 100% (3370/3370), done.
#Total 3370 (delta 1333), reused 2891 (delta 1108)
du -sh .git/*
#7.5M	.git
git gc --aggressive --prune=now
#Counting objects: 3370, done.
#Delta compression using up to 4 threads.
#Compressing objects: 100% (3155/3155), done.
#Writing objects: 100% (3370/3370), done.
#Total 3370 (delta 1419), reused 1944 (delta 0)
du -sh .git
#7.3M	.git
```

