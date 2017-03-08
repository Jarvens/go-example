# goLang
## 什么是Go ##
###Go编程语言是一个使得程序员更加有效率的开源项目。Go是具有表达力、简洁、清晰和有效率的。###
###它的并行机制使其很容易编写多核和网络应用,而新奇的类型系统允许构建有弹性的模块化程序。###
###GO编译到机器码非常快速，同时具有遍历的垃圾回收和强大的运行时反射。###
###它是快速的、静态类型编译语言，但是感觉上是动态类型的，解释性语言。###

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



### Chanpter-2 变量 基本类型

- Go语言同其它语言不通的地方在于变量的类型在变量名后面。例如JAVA:  int a;  Go:  a int;
- 当定义了一个变量,它默认值为null值 例如 a int;  a=0  || a string;  a=""
- 在Go中变量的声明和赋值是两个过程,但是可以连在一起,例如

```
    var a int;
    var b int;
    a=1;
    b=2;
    fmt.Println(a+b);
```
- :=声明并赋值的方式仅用于函数体内
- 多类型声明,注意小括号的用法,在Go语言中,声明但未使用的变量在运行时会报错
```
    var (
        c int
        d bool
    )

    fmt.Println(c,d);
```
- 相同类型多个变量的声明方式
```
  var e,f int;
  e,f = 1,2;
  fmt.Println(e,f);
  g,h:=1,2;
  fmt.Println(g,h);
```

- 特殊变量 _,此时i 为 4 ,如果打印_的话,控制台会输出cannot use _ as value ,究其原因为:任何值赋给_都将会被丢弃
```
   _,i := 3,4;
   fmt.Println(i);
```

### Chapter-3 int
- Go的 int类型 是由硬件决定适当的长度,意味着在32位硬件上 是32的 在64位的硬件上是64位的,注意()
- 如果你希望明确其长度,可以使用int32 或者 uint 32 完整的证书类型列表(符号和无符号)是 int8,int16,int32,int64 和byte,unit8,uint16,uint32,uint64
- byte是uint8的别名 浮点类型有 float32和float64(没有float类型),64位的整数和浮点数总是64位的,即使在32位的架构上,需要注意的是这些类型全部是独立的,并且混合使用这些类型向变量赋值会引起编译器错误
```
    //引起编译错误   invalid operation: a + b (mismatched types int and int32)
    var a int;
    var b int32;
    a = 10;
    b = 12;
    fmt.Println(a + b);
```
```
   const(
        c =iota
        d = iota
    )
    fmt.Println(c,d);
```
```
   const(
        e =iota
        f      //隐士调用iota
    )

    fmt.Println(e,f);
```
```
    var g string ="中国";
    var h float32 = 3.1415926;
    fmt.Println(g,h);
```


### Chapter-4 字符 & 字符串

- rune是int32的别名,用UTF-8进行编码。这个类型使用的情景为:例如需要遍历字符串中的字符,可以循环每个字节,隐刺为了获得实际的字符,需要使用rune类型
- Go语言中是UTF-8编码的,字符串是由双引号  包裹 "" 的字符序列,如果是用单引号 '' 则表示一个字符(UTF-8编码)这种在Go中不是string.一旦给变量赋值就不可更改,在Go中字符串是不可变的
(因为字符串从创建到销毁仅指向一个内存地址,如果字符串改变,那么指向的内存地址就不是原来的地址,所以是说不可变)
- 示例:定义 str为字符串  将字符串转换为数组并且将feture的第一个字符替换为大写的H 最后强制转换为字符串并且赋值给str1

```
    str := "hello,world";
    feture := []rune(str);
    feture[0] = 'H';
    str1 := string(feture);
    fmt.Println(str1);
```

### 








