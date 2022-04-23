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

    // 正の無限大と負の無限大とNaN
    zero := 0.0
    pinf := 1.0 / zero
    ninf := -1.0 / zero
    nan := 0.0 /zero
    fmt.Printf("pinf = %v, ninf = %v, nan = %v\n", pinf, ninf, nan)

    // 浮動小数点型から整数型へのキャストは切り捨てになる
    fnum := 2.52
    fmt.Printf("fnum = %f, int(fnum) = %d\n", fnum, int(fnum))

    // rune型
    r := '曜'
    fmt.Printf("r = '曜' = %v\n", r)

    // 文字列とRAW文字列リテラル
    str := "千歌"
    rawlit := `よう
ちか
りこ
`
    fmt.Printf("str = %s\n", str)
    fmt.Printf("%s", rawlit)

    // 配列
    intarr := [9]int{1,2,3,4,5,6,7,9}
    fmt.Printf("intarr = %v\n", intarr)
    // intの場合は初期値を入れないと0で初期化される
    var intarr2 [3]int
    fmt.Printf("intarr2 = %v\n", intarr2)
    // 他のデータ型でも自動で初期化される
    var strarr [3]string
    fmt.Printf("strarr = %v\n", strarr)
    var boolarr [3]bool
    fmt.Printf("boolarr = %v\n", boolarr)
    // 初期値を指定する場合は要素数を指定しない書き方もできる
    intarr3 := [...]int{445, 745, 25, 874}
    fmt.Printf("intarr3 = %v\n", intarr3)

    // 配列の代入
    strarr1 := [3]string{"honoka", "kotori", "umi"}
    strarr2 := [3]string{"ayumu", "setsuna", "ai"}
    strarr1 = strarr2
    // 配列の代入はコピーとなるため、この場合strarr2[1]に代入してもstrarr1[1]には影響しない
    strarr2[1] = "kasumi"
    fmt.Printf("strarr1 = %v, strarr2 = %v\n", strarr1, strarr2)

    // interface型
    var itfc interface{}
    fmt.Printf("itfc = %v\n", itfc)
    // どのデータ型でも代入できる
    itfc = int(32)
    itfc = "kuku"
    itfc = false
}
