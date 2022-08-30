package main

import (
	"fmt"
)

type User struct {
	Name string
	Age int
}

func main() {
	// fmt.Printf()で%を表示する場合は%%を使用する
	fmt.Printf("%%\n")

	// 指定した桁数で右詰め
	fmt.Printf("%5d\n", 445)
	// 指定した桁数で0埋め
	fmt.Printf("%05d\n", 445)
	// 10進数 -> 2進数
	fmt.Printf("%b\n", 445)

	// 小数点以下を指定した値で丸める
	fmt.Printf("%.3f\n", 1.8742525)
	// 小数点以下を指定した桁で丸め、全体を指定した桁で右詰め
	fmt.Printf("\"%10.3f\"\n", 1.8742525)
	// 上記の左詰め版
	fmt.Printf("\"%-10.3f\"\n", 1.8742525)

	// 文字数で右詰め
	fmt.Printf("\"%10s\"\n", "南ことり")
	// 文字数で左詰め
	fmt.Printf("\"%10s\"\n", "南ことり")
	// ""付き
	fmt.Printf("\"%q\"\n", "南ことり")

	// bool型専用
	fmt.Printf("%t\n", true)
	// ポインタ型専用
	you := "You Watanabe"
	fmt.Printf("%p\n", &you)
	// データ型
	fmt.Printf("%T\n", &you)

	u1 := &User {
		Name: "Chika Takami",
		Age: 16,
	}
	// %vは様々なデータ型に対応
	fmt.Printf("%v\n", u1)
	// %+vの場合は構造体のフィールド名も含まれる
	fmt.Printf("%+v\n", u1)
	// %#vの場合はデータ型も含まれる
	fmt.Printf("%#v\n", u1)
}
