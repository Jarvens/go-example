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

    sum := 0
}

