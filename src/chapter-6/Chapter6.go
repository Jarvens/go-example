package main;

import (
    "fmt"
    "unicode/utf8"
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

func loop_func() int {
    loop_counter := 0;
    Loop:
    if loop_counter < 10 {
        loop_counter += 1;
        goto Loop;
    }
    return loop_counter;
}

/**
递归函数
 */

func rec(i int) {
    if i == 10 {
        return
    }
    rec(i + 1);
    fmt.Printf("%d ", i)
}

func currrent(x, y int) (a1, b1 int) {
    a1 = x + y
    b1 = x - y
    return
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
    array2 := [...][5]int{[5]int{11, 12, 13, 14, 15}, [5]int{16, 17, 18, 19, 20}}
    for _, value := range array2 {
        for _, value1 := range value {
            fmt.Println(value1);
        }
    }

    //创建array数组
    array3 := [...]int{9, 8, 7, 6, 5, 4, 3, 2, 1}

    //打印数组长度
    fmt.Println(len(array3));

    //打印数组序号
    fmt.Println(cap(array3));

    //创建slice ,slice是一个指向底层array 的指针  array[2:5] 代表从array  取出 n ~ m-1 序号的字符串
    sl := array3[2:5];
    fmt.Println(sl);

    //将array全部赋值给 slice
    sl1 := array3[:];
    fmt.Println(sl1);

    //将array 从0 ~ m-1 赋值给slice
    sl2 := array3[:6];
    fmt.Println(sl2);

    //数组追加
    var sl3[] int;
    sl4 := append(sl3, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0);
    fmt.Println(sl4);

    //仅仅声明slice
    sl5 := make([]int, 10);
    fmt.Println(len(sl5));
    fmt.Println(cap(sl5));

    //声明map
    map_source := make(map[string]int);
    fmt.Println("map_source初始长度为 ->", len(map_source));

    //声明并赋值   key,value
    map_month := map[string]int{
        "Jan":31, "Feb":28, "Mar":31, "Apr":30, "May":31,
    }
    fmt.Println(map_month["Jan"]);
    //向map追加元素
    map_month["Jun"] = 30;
    fmt.Println("新mao_month ->", map_month);

    //检测值是否存在
    value, flag := map_month["Dec"];
    fmt.Println(value, flag);

    //Q1 练习题
    loop_counter := 0;
    for i := 0; i < 10; i++ {
        loop_counter += 1;
    }
    fmt.Println("计数器为 ->", loop_counter);

    fmt.Println(loop_func())

    //Q2 练习  3 5 倍数
    const (
        FIZZ = 3
        BUZZ = 5
    )
    var p bool = false;
    for i := 1; i < 100; i++ {
        if i % FIZZ == 0 && i % BUZZ == 0 {
            fmt.Println("FIZZBUZZ ->", i);
            p = true;
        }
        if i % BUZZ == 0 {
            fmt.Println("5的倍数 ->", i);
            p = true;
        }
        if !p {
            fmt.Println("啥都不是 ->", i);
        }
        fmt.Println()
    }

    str := "A";
    for i := 0; i < 100; i++ {
        fmt.Println(str);
        str = str + "A";
    }

    str1 := "dsjkdshdjsdh....js"
    fmt.Println(len([]byte(str1)))
    fmt.Println(utf8.RuneCount([]byte(str1)));

    s_value := "ewlgkmerithyeiltun";
    str2 := []rune(s_value);
    //copy(str2[4:7],[]rune("ABC"))
    fmt.Printf("After : %s\n", string(str2));
    fmt.Println(string(str2));
    rec(0);

    fmt.Println(currrent(1,2))

    values:= func(arg ...int) {
        for _,value :=range arg{
            fmt.Println(value)
        }
    }

    values(2,6,1,4,9,2)

}

