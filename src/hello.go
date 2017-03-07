package main

import "fmt";

func main() {
    s := "helloworld";
    c:=[]rune(s);
    c[0]='H';
    s1:=string(c);
    fmt.Printf(s1);
}
