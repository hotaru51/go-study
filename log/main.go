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

	// ログのフォーマットを指定する
	sampleText := "log test"
	// 標準
	log.SetFlags(log.LstdFlags)
	log.Println(sampleText)
	// マイクロ秒追加
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	log.Println(sampleText)
	// 時刻とファイルの行番号
	log.SetFlags(log.Ltime | log.Lshortfile)
	log.Println(sampleText)

	log.SetFlags(log.LstdFlags)
	// プレフィックスを設定する
	log.SetPrefix("[log] ")
	log.Println(sampleText)

	// ロガーの生成
	logger := log.New(os.Stdout, "[log] ", log.LstdFlags)
	logger.Println("my logger")
}
