package main

import (
	"fmt"
	"runtime"
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

	fmt.Printf("NumCPU: %d\n", runtime.NumCPU())
	fmt.Printf("NumGoroutine: %d\n", runtime.NumGoroutine())
	fmt.Printf("Versions: %s\n", runtime.Version())
}
