package main;

import "fmt";
func swap(x,y string)(string){
    return x,y;
}
func main() {
    /**
    单一类型声明
     */
    var a int;
    var b int;
    a=1;
    b=2;
    fmt.Println(a+b);

    /**
    多类型声明
     */
    var (
        c int
        d bool
    )

    fmt.Println(c,d);

    /***
    多个类型相同的变量声明
     */
    var e,f int;
    e,f = 1,2;
    fmt.Println(e,f);

    g,h:=1,2;
    fmt.Println(g,h);

    _,i := 3,4;
    fmt.Println(i);
    j:= swap("hello","world");
    fmt.Println(j);

}
