package main

import (
	"fmt"
	"os"
	// "log"
)

func openFile(path string) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	// deferで開いたファイルを確実にcloseする
	fmt.Println(f.Name())
	defer f.Close()
}

func main() {
	// os.Exitで終了するとdeferは実行されない
	defer func() {
		fmt.Println("defer")
	}()

	// 存在しないファイルを開こうとした場合にエラーメッセージを出力
	_, err := os.Open("hoge")
	if err != nil {
		// log.Fatal(err)
	}

	// コマンドライン引数はos.Args[]で参照できる
	for i, v := range os.Args {
		fmt.Printf("os.Args[%d] = \"%s\"\n", i, v)
	}

	openFile("./uranohoshi.txt")

	os.Exit(0)
}
