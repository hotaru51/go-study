package main

import (
	"fmt"
	"os"
	"log"
)

func main() {
	// os.Exitで終了するとdeferは実行されない
	defer func() {
		fmt.Println("defer")
	}()

	_, err := os.Open("hoge")
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}
