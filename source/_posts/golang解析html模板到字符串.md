title: golang解析html模板到字符串
date: 2015-01-30 00:05:22
tags: golang
categories:
- 术业
- golang
---
&emsp;&emsp;说来自己对golang的掌握也就半壶水,不过我向来认为学习就是一个不断踩坑的过程.只有在实践中才能不断的取得进步.最近在独力编写golang版本的影梭管理面板,其中有个需求就是给用户发送激活邮件,本着尽善尽美的原则,计划提供可编辑的html邮件模板.那么,问题就来了.使用`html/template`包对html模板进行解析,但是`Execute`函数是输出到`io.Writer`,而我需要的是把模板解析结果保存到`string`类型的变量中,以便直接传送给下一步负责发送邮件的函数使用.<!-- more -->
&emsp;&emsp;在网络上查阅资料后得知实现上述要求方法非常简单,只需要设置一个`bytes.Buffer`,`Execute`的时候把数据写入到这个缓存器当中,再调用`String`方法转换成字符串即可.
&emsp;&emsp;Don't talk,show me code!So,there is a sample:

``` golang
package main

import (
    "html/template"
    "bytes"
    "fmt"
)

type Mail struct {
    UserName string
    SiteName string
    ActiveLink string
}

func main() {
    var doc bytes.Buffer
    var templateString = "{&#123;.UserName}}你好,您在{&#123;.SiteName}}注册了帐号,请点<a href=\"{&#123;.ActiveLink}}\">击这里激活!</a>"
    t := template.New("fieldname example")
    t, _ = t.Parse(templateString)
    p := Mail{UserName: "Frank",SiteName: "SSP", ActiveLink: "http://xizhimen.com/active/9a32f2"}
    t.Execute(&doc,p)
    html := doc.String()
    fmt.Println(html)
}
```
&emsp;&emsp;Sweet dream,Frank.


参考文献:

[1] [Package template](http://golang.org/pkg/html/template/)

[2] [Package bytes #Buffer.string](http://golang.org/pkg/bytes/#Buffer.String)

[3] [Make template.Execute write output to a string](https://groups.google.com/forum/#!topic/golang-nuts/dSFHCV-e6Nw)