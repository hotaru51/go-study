package main

import (
	"fmt"
	"bufio"
	"strings"
)

func main() {
	s := `honoka kotori umi
hanayo rin maki
nico eri nozomi`

	r := strings.NewReader(s)
	scanner := bufio.NewScanner(r)

	// Scannerのscanメソッドを変更する
	// 単語毎に読み取る
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
