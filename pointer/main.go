package main

import (
	"fmt"
)

func main() {
	// ポインタ型の初期値はnil
	var p *int
	fmt.Printf("(p == nil) = %v\n", (p == nil))

	// 変数のアドレス取得
	i := 5
	pi := &i
	fmt.Printf("pi => %T\n", pi)
}
