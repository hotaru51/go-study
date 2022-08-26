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
}
