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

	// float -> string
	f := float64(445.745)
	strf1 := strconv.FormatFloat(f, 'f', 2, 64)
	fmt.Printf("strf1 = %s\n", strf1)
	strf2 := strconv.FormatFloat(f, 'g', -1, 64)
	fmt.Printf("strf2 = %s\n", strf2)

	// string -> bool
	tList := []string{"true", "1", "t", "T", "TRUE", "True"}
	for _, str := range(tList) {
		b, err := strconv.ParseBool(str)
		fmt.Printf("str, b, err = %s, %t, %v\n", str, b, err)
	}
	fList := []string{"false", "0", "f", "F", "FALSE", "False"}
	for _, str := range(fList) {
		b, err := strconv.ParseBool(str)
		fmt.Printf("str, b, err = %s, %t, %v\n", str, b, err)
	}
}
