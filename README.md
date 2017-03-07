# goLang
### 什么是Go ###
Go编程语言是一个使得程序员更加有效率的开源项目。Go是具有表达力、简洁、清晰和有效率的。
它的并行机制使其很容易编写多核和网络应用,而新奇的类型系统允许构建有弹性的模块化程序。
GO编译到机器码非常快速，同时具有遍历的垃圾回收和强大的运行时反射。
它是快速的、静态类型编译语言，但是感觉上是动态类型的，解释性语言。
### 环境配置 ###
- 从官网[http://www.golangtc.com/]下载Go语言安装包
- 配置环境变量,以下是本机的环境变量配置信息
当前home目录下的.bash_profile
```
export GOPATH=/Users/kunlun/Documents/WebStorm/go-example
export GOBIN=$GOPATH/bin
export PATH=$PATH:/$GOBIN
```
etc目录下profile文件
```
export GOROOT=/usr/local/go
export PATH=$PATH:$GOROOT
```

### Chapter-1:HelloWorld ###
- 知识点
- 所有go文件以 ```package xxx``` 开头,对于独立运行的执行文件必须是 ```package main```
- 接下来是```import xxx```,在本次示例中我们引入了fmt包的函数用来打印字符串,在Go语言中字符串由 " 包裹,并且可以包含非ASCll的字符
- 编译&运行:
- 构建HelloWorld只需要在控制台执行  ```go build HelloWorld.go```,编译后在工作目录下得到一个可执行文件 HelloWorld
