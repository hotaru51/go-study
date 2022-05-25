package main

import (
	"fmt"
)

func main() {
	// deferはpanic後でも実行される
	defer fmt.Println("kotori")
	panic("runtime error")
	fmt.Println("honoka")
}
