package main

import (
	"fmt"
	"unicode"
)

func main() {
	// 数値を判別
	fmt.Printf("IsDigit(\"X\") = %t\n", unicode.IsDigit('X'))
	fmt.Printf("IsDigit(\"3\") = %t\n", unicode.IsDigit('3'))
	fmt.Printf("IsDigit(\"３\") = %t\n", unicode.IsDigit('３'))
	fmt.Printf("IsDigit(\"三\") = %t\n", unicode.IsDigit('三'))

	// 文字を判別
	fmt.Printf("IsLetter(\"A\") = %t\n", unicode.IsLetter('A'))
	fmt.Printf("IsLetter(\"Ａ\") = %t\n", unicode.IsLetter('Ａ'))
	fmt.Printf("IsLetter(\"3\") = %t\n", unicode.IsLetter('3'))
	fmt.Printf("IsLetter(\"３\") = %t\n", unicode.IsLetter('３'))
	fmt.Printf("IsLetter(\"あ\") = %t\n", unicode.IsLetter('あ'))
	fmt.Printf("IsLetter(\"♨\") = %t\n", unicode.IsLetter('♨'))

	// スペースを判別
	fmt.Printf("IsSpace(\" \") = %t\n", unicode.IsSpace(' '))
	fmt.Printf("IsSpace(\"\\t\") = %t\n", unicode.IsSpace('\t'))
	fmt.Printf("IsSpace(\"　\") = %t\n", unicode.IsSpace('　'))
	fmt.Printf("IsSpace(\"_\") = %t\n", unicode.IsSpace('_'))

	// 制御文字
	fmt.Printf("IsControl(\"\\n\") = %t\n", unicode.IsControl('\n'))
}
