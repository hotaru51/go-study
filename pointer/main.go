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

	// 関数への参照渡し
	i2 := 1
	fmt.Printf("i2 = %d\n", i2)
	inc(&i2)
	fmt.Printf("i2 = %d\n", i2)
	inc(&i2)
	fmt.Printf("i2 = %d\n", i2)
	inc(&i2)
	fmt.Printf("i2 = %d\n", i2)
}

func inc(i *int) {
	*i++
}
