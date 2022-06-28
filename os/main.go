package main

import (
	"fmt"
	"os"
)

func main() {
	// os.Exitで終了するとdeferは実行されない
	defer func() {
		fmt.Println("defer")
	}()

	os.Exit(0)
}
