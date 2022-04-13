package main

import (
    "fmt"
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
}
