package main

import (
	"log"
	"os"
)

func main() {
	log.Print("aaaaa\n")
	log.Println("bbbbb")
	log.Printf("%s\n", "ccccc")

	// デフォルトでは標準エラー出力になるので、標準出力に変更する
	log.SetOutput(os.Stdout)
	log.Println("stdout")
}
