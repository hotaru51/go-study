package main

import (
	"fmt"
)

func main() {
	// deferはpanic後でも実行される
	defer func() {
		if x := recover(); x != nil {
			fmt.Printf("recover: %s", x)
		}
	}()
	panic("runtime error")
	fmt.Println("honoka")
}
