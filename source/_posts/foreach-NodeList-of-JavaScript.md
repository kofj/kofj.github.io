---
title: JS 遍历 NodeList 对象
tags:
  - JavaScript
categories:
  - 术业
  - 前端
date: 2016-12-14 04:06:52
---

## 0x1 问题描述
&emsp;&emsp;本博客使用的 `Yilia` 主题在移动设备上有这样一个问题：标签无法正常显示。使用 Chrome 浏览器进行远程调试，捕捉到了下列错误信息：
```JavaScript
Uncaught TypeError: $tags.forEach is not a function
init @ fix.js?9699:25
window.onload @ main.js?8359:24
```
<!-- more -->
## 0x2 溯源
&emsp;&emsp;JavaScript 的继承机制是基于原型的。例如，一个数组元素 `arr` 上之所以有一些数组方法（比如 `forEach`）,是因为它的原型链上有这些方法：
&emsp;&emsp;arr → Arrary.prototype → Object.prototype → null
&emsp;&emsp;`NodeList` 元素 tags 的原型链如下：
&emsp;&emsp;tags → NodeList.prototype → Object.prototype → null
&emsp;&emsp;如果 NodeList 的原型上没有 `forEach` 方法，调用该方法必然会报错。查阅文档可知，[DOM4](https://www.w3.org/TR/dom/#interface-nodelist) 中的 `NodeList` 接口只实现了 `item` 方法。
&emsp;&emsp;但是，只是一些旧的浏览器中 `NodeList` 对象还没有实现 `forEach`、`values` 这些方法。所以，尽管 `NodeList` 不是 `Array`，它也是能够使用 `forEach` 方法遍历元素的。
&emsp;&emsp;我在手机上使用的 Chrome 浏览器比较老旧，版本是 `47.0.2526.83`。使用 `Object.getPrototypeOf` 查看 `NodeList` 原型，发现确实是只实现了 `item` 方法。而在 Macbook 上安装的则是最新版的 Chrome 浏览器。
&emsp;&emsp;归根结底，问题原因是处在了对老版本浏览器的兼容上了。
## 0x3 解决方案
&emsp;&emsp;我们可以简单的通过 `prototypes` 扩展 NodeList 的方法来兼容老浏览器。不过这样做可能会遇到一些问题，具体可以查阅参考文献第三条。
```JavaScript
NodeList.prototype.forEach = Array.prototype.forEach
```
&emsp;&emsp;除了通过 `prototypes` 扩展 `DOM` 方法外，我们还可以直接使用 `for` 循环或者是自定义一个 forEach 的函数来遍历 NodeList。

## 参考文献
1. [NodeList](https://developer.mozilla.org/en-US/docs/Web/API/NodeList), https://developer.mozilla.org/en-US/docs/Web/API/NodeList
1. [DOM4](https://www.w3.org/TR/dom/#interface-nodelist), https://www.w3.org/TR/dom/#interface-nodelist
1. [Ditch the [].forEach.call(NodeList) hack](https://toddmotto.com/ditch-the-array-foreach-call-nodelist-hack/), https://toddmotto.com/ditch-the-array-foreach-call-nodelist-hack/
