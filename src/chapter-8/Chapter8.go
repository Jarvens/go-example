package main;

import (
    "fmt"
)

func main() {

    // p 为指向 int 类型的指针,并且值为nil,也就是null值 又称为空指针
    var p *int;
    fmt.Println(p)

    // &  为取址操作符  整体含义为  定义i为整形变量,并定义j为指向i地址的数据
    var i_1 int = 15;
    var j = &i_1;
    fmt.Println(i_1)
    fmt.Printf("%d\n",j)
    for i:=0;i<10 ;i++  {
        i_1++;
    }
    fmt.Println(i_1)
    fmt.Println(j)

}
