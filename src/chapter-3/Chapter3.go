package main;

import "fmt";

func main() {
    /**
    混合数据类型计算  报错   invalid operation: a + b (mismatched types int and int32)
     */
    var a int;
    var b int32;
    a = 10;
    b = 12;
    fmt.Println(a + b);
    /**
    常量类型  const  以下两种写法是相等的iota 初始值为0
     */
    const(
        c =iota
        d = iota
    )
    fmt.Println(c,d);

    const(
        e =iota
        f      //隐士调用iota
    )

    fmt.Println(e,f);

    /**
    省略  var关键字只能在声明常量的时候省略
     */

    var g string ="中国";
    var h float32 = 3.1415926;
    fmt.Println(g,h);
}
