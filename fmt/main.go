package main

import (
	"fmt"
)

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
}
