package main;

import (
    "fmt"
    "math/cmplx"
    "math"
);

func test(sum int) (x, y int) {
    x = sum / 4;
    y = sum % 3;
    return
}
func main() {

    /**
    测试函数
     */
    fmt.Println("Now you have %g problems.", math.Pi);
    fmt.Println(test(28));

    var (
        a bool = false
        maxb uint64 = 1 << 64 - 1
        c complex128 = cmplx.Sqrt(-5 + 12i)
    )
    const f = "%T(%v)\n";
    fmt.Println(f, a, a);
    fmt.Println(f, maxb, maxb);
    fmt.Println(f, c, c);

    var a1 int;
    var a2 float64;
    var a3 string;
    var a4 bool;
    fmt.Println(a1,a2,a3,a4);
}

