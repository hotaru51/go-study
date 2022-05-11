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
	nan := 0.0 / zero
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
	intarr := [9]int{1, 2, 3, 4, 5, 6, 7, 9}
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

	// 関数
	fmt.Printf("add(1 + 2) = %d\n", add(1, 2))
	say("ohayohane")
	q1, r1 := div(10, 3)
	fmt.Printf("div(10, 3) = %d 余り %d\n", q1, r1)

	// 戻り値の破棄
	_, r2 := div(7, 3)
	fmt.Printf("div(7, 3)の余り = %d\n", r2)

	fmt.Printf("retvar() = %d\n", retvar())

	// 無名関数
	fn1 := func(x, y int) int { return x + y }
	fmt.Printf("%T\n", fn1)
	fmt.Printf("f(3, 2) = %d\n", fn1(3, 2))
	fmt.Printf("%#v\n", func(x, y int) int { return x * y })
	fmt.Printf("returnFunction() -> %#v, (returnFunction())() -> %s\n", returnFunction(), (returnFunction())())

	// クロージャ
	fn2 := later()

	fmt.Println(fn2("nocchi"))
	fmt.Println(fn2("kashiyuka"))
	fmt.Println(fn2("a-chan"))
	fmt.Println(fn2("yskt"))

	// 定数
	const X = 2

	// 定数の値省略
	// 全て同じ値になる
	const (
		CHIKA = 1
		YOU
		RIKO
	)
	fmt.Printf("CHIKA = %d, YOU = %d, RIKO = %d\n", CHIKA, YOU, RIKO)

	// 型あり定数
	const I64 = int64(10)
	fmt.Printf("I64 = %d (%T)\n", I64, I64)

	// enum型っぽいやつ
	// iotaは呼び出すたびに1ずつ増分する
	const (
		A = iota
		B
		C
	)
	fmt.Printf("A = %d, B = %d, C = %d\n", A, B, C)

	// for
	// 下記は無限ループ
	/*
	   for {
	       fmt.Println("逮捕")
	   }
	*/

	// いつもの
	member1 := [...]string{"chika", "riko", "you"}
	for i := 0; i < 3; i++ {
		fmt.Println(member1[i])
	}

	// if
	m1 := 1
	// m1 := 2
	// m1 := 3
	if m1 == 1 {
		fmt.Println("m1 == 1")
	} else if m1 == 2 {
		fmt.Println("m1 == 2")
	} else {
		fmt.Println("m1 != 1 && m1 != 2")
	}

	// 簡易文付きif
	// ここで定義した変数はブロック内のみ有効
	if x, y := 1, 2; x < y {
		fmt.Printf("%d < %d\n", x, y)
	}

	// 他の言語のwhileの様なfor
	i = 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	// rangeを使用したfor
	for i, s := range member1 {
		fmt.Printf("member1[%d] = %s\n", i, s)
	}

	// stringに対してrangeを使用するとrune型が返る
	fmt.Printf("name = %s(%T)\n", name, name)
	for i, r := range name {
		fmt.Printf("name[%d] = %v(%T) = %s\n", i, r, r, string(r))
	}

	// switch
	sw1 := 3
	switch sw1 {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("unknown")
	}

	// CやJavaのswitchのようにフォールスルーさせる
	sw2 := "N"
	switch sw2 {
	case "N":
		sw2 += "I"
		fallthrough
	case "NI":
		sw2 += "C"
		fallthrough
	default:
		sw2 += "O"
	}
	fmt.Println(sw2)

	// 簡易分付きswitch
	// if同様、下記の変数はswitch内でのみ有効
	switch member := "you"; member {
	case "yoshiko", "ruby", "hanamaru":
		fmt.Printf("%s = 1年生\n", member)
	case "chika", "you", "riko":
		fmt.Printf("%s = 2年生\n", member)
	case "kanan", "mari", "dia":
		fmt.Printf("%s = 3年生\n", member)
	default:
		fmt.Println("unknown")
	}

	// 型アサーション
	var if1 interface{} = 3
	chk1 := if1.(int)     // chk1にはint型で3が入る
	// chk2 := if1.(float64) // これは実行時エラー
	fmt.Printf("%d(%T)\n", chk1, chk1)
	// fmt.Println(chk1, chk2)
}

// int xとyを足して返す
func add(x int, y int) int {
	return x + y
}

// 戻り値がない場合
func say(msg string) {
	fmt.Printf("%s\n", msg)
}

// 複数の戻り値を返す
func div(x, y int) (int, int) {
	q := x / y
	r := x % y

	return q, r
}

// 戻り値を表す変数
// この関数はint(1)を返す
func retvar() (a int) {
	a = 1

	return
}

// 関数を返す関数
func returnFunction() func() string {
	return func() string { return "sakurauchi" }
}

// クロージャ
func later() func(string) string {
	var store string

	return func(next string) string {
		s := store
		store = next

		return s
	}
}
