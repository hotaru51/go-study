package main

import (
    "fmt"
)

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
}
