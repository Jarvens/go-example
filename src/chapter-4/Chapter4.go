package main;

import "fmt";

func main() {

    /**
    字符串替换示例
     */
    str := "hello,world";
    feture := []rune(str);
    feture[0] = 'H';
    str1 := string(feture);
    fmt.Println(str1);

    /**
    错误示例
     */

    a := "hello"
    + "world";
    fmt.Println(a);

    /**
    正确示例
     */

    b := "hello" + "world";
    fmt.Println(b);
}
