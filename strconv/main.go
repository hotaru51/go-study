package main

import (
	"fmt"
	"strconv"
)

func main() {
	// bool -> string
	b := true
	strb := strconv.FormatBool(b)
	fmt.Printf("strb = %s\n", strb)

	// int -> string
	i := int64(-445)
	// 第2引数は基数
	stri := strconv.FormatInt(i, 10)
	fmt.Printf("stri = %s\n", stri)
}
