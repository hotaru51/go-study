package main

import (
    "fmt"
    "math"
)

// パッケージ変数
var pkg = "aaa"

func main() {
    // fmt.Println
    // 複数渡すとスペース区切りで出力
    fmt.Println("auie")
    fmt.Println("nocchi", "kashiyuka", "a-chan")

    // fmt.Printf
    fmt.Printf("a = %d\n", 5)

    // %v, %#v, %T
    fmt.Printf("数値=%v, 文字列=%v, 配列=%v\n", 417, "渡辺曜", [...]int{4, 1, 7})
    fmt.Printf("数値=%#v, 文字列=%#v, 配列=%#v\n", 417, "渡辺曜", [...]int{4, 1, 7})
    fmt.Printf("数値=%T, 文字列=%T, 配列=%T\n", 417, "渡辺曜", [...]int{4, 1, 7})

    // Print, Println
    print("hanamaru ")
    print("ruby ")
    print("yoshiko ")
    println("yohane")
    println("dia kanan mari")

    // 変数定義
    var n int
    var x, y, z int
    var (
        a, b int
        name string
    )

    // 代入
    n = 3
    x, y, z = 4, 4, 5
    a, b = 2, 5
    name = "kotori"

    fmt.Printf("n = %d\n", n)
    fmt.Printf("x, y, z = %d, %d, %d\n", x, y, z)
    fmt.Printf("a, b = %d, %d\n", a, b)
    fmt.Printf("name = %s\n", name)

    // := で定義と代入を同時に行う、型推論
    i := 445
    // i := 874 <- := -で再代入は不可
    f := 7.45
    s := "honoka"
    fmt.Printf("%d %f %s\n", i, f, s)

    var (
        nico = 3
        maki = 1
        eli  = 3
    )
    fmt.Printf("nico = %d , maki = %d, eli = %d\n", nico, maki, eli)
    fmt.Println(pkg)

    /*
     * データ型
     */

    // bool
    var bb bool
    bb = true
    fmt.Printf("b = %v\n", bb)

    // 整数型のキャスト
    var_int := 1
    var_byte := byte(var_int)
    var_i64 := int64(var_int)
    fmt.Printf("var_int = %d, var_byte = %d, var_i64 = %d\n", var_int, var_byte, var_i64)

    // mathパッケージで各型の最大値を出す
    fmt.Printf("int32_max = %d\n", math.MaxInt32)
    fmt.Printf("int64_max = %d\n", math.MaxInt64)
    fmt.Printf("int64_min = %d\n", math.MinInt64)

    // 浮動小数点型
    fmt.Printf("float32_max = %f\n", math.MaxFloat32)
    fmt.Printf("float64_max = %f\n", math.MaxFloat64)
    fmt.Printf("float32_smallest_nonzero = %f\n", math.SmallestNonzeroFloat32)
    fmt.Printf("float64_smallest_nonzero = %f\n", math.SmallestNonzeroFloat64)

    // 暗黙的な変数の代入ではfloat64となる
    f64 := 0.874
    fmt.Printf("f64 = %T\n", f64)
}
