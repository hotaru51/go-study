package main

import (
	"fmt"
)

func sub() {
	for i := 0; i < 10; i++ {
		fmt.Println("chika")
	}
}

func main() {
	// ゴルーチン開始
	go sub()
	// 無名関数も指定できる
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("riko")
		}
	}()
	for i := 0; i < 10; i++ {
		fmt.Println("you")
	}
}
