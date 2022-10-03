package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 標準入力を読み込むスキャナを生成
	scanner := bufio.NewScanner(os.Stdin)

	// 入力の読み込みが成功するまでループ
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	// 読み込みが失敗した場合の処理
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "読み込みエラー:", err)
	}
}
