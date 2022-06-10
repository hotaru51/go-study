package main

import (
	"fmt"
)

func main() {
	// ポインタ型の初期値はnil
	var p *int
	fmt.Printf("(p == nil) = %v", (p == nil))
}
