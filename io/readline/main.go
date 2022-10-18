package main

import (
	"fmt"
	"bufio"
	"strings"
)

func main() {
	s := `honoka
kotori
umi
`
	// 文字列からReader生成
	r := strings.NewReader(s)

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
