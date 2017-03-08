package main;

import (
    "fmt"
);

func compare(a, b[]byte) int {
    for i := 0; i < len(a)&&i < len(b); i++ {
        switch {
        case a[i] > b[i]:
            return 1;
        case a[i] < b[i]:
            return -1;
        }
    }
    return 0
}

func main() {
    /**
    for  break  跳出 i循环
     */
    I:  for i := 0; i < 5; i++ {
        for j := 0; j < 10; j++ {
            if j > 8 {
                break I;
            }
            fmt.Println(j);
        }
    }

    /**
    for continue  跳出本次循环进入下一循环
     */
    for k := 0; k < 10; k++ {
        if k > 5 {
            continue
        }
        fmt.Println(k);
    }

    /**
    range 循环
     */
    list := []string{"a", "b", "c", "d", "e", "f"};
    for k, v := range list {
        fmt.Println(k, v);
    }

    /**
    range  直接作用在字符串上
     */
    for key, value := range "defghijklmn" {
        fmt.Println(key, value);
    }

    charater := 2;
    switch charater {
    case 0:
        fmt.Println("零");
    case 1:
        fmt.Println("一");
    default:
        fmt.Println("不匹配");
    }

    //定义一维数组
    array := [...]int{1, 2, 3};
    for _, value := range array {
        fmt.Println(value);
    }

    //定义二维数组 2个数组 每个数组长度为3
    array1 := [2][3]int{[...]int{1, 2, 3}, [...]int{4, 5, 6}}
    for _, value1 := range array1 {
        for _, value := range value1 {
            fmt.Println(value);
        }
    }

    //二维数组简写
    array2 := [...][5]int{[5]int{11, 12, 13, 14, 15}, [5]int{16, 17, 18, 19,20}}
    for _, value := range array2 {
        for _, value1 := range value {
            fmt.Println(value1);
        }
    }
}

