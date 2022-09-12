package main

import (
	"fmt"
	"strings"
)

func main() {
	// []string内の要素を結合する
	cyaron := []string{"you", "chika", "ruby"}
	fmt.Println(strings.Join(cyaron, "! ") + "!")

	// 文字列の検索
	fmt.Printf("Index(\"ABCDE\", \"A\") = %d\n", strings.Index("ABCDE", "A"))
	fmt.Printf("Index(\"ABCDE\", \"BCD\") = %d\n", strings.Index("ABCDE", "BCD"))
	fmt.Printf("Index(\"ABCDE\", \"Y\") = %d\n", strings.Index("ABCDE", "Y"))

	// 最後に文字列が見つかる部分を検索
	fmt.Printf("LastIndex(\"ABCDEABCD\", \"B\") = %d\n", strings.LastIndex("ABCDEABCD", "B"))

	// 指定した文字列のいずれかの文字が最初に見つかる部分を返す
	fmt.Printf("IndexAny(\"ABCDEABCD\", \"CDE\") = %d\n", strings.IndexAny("ABCDEABCD", "CDE"))
	//IndexAnyの最後に見つかる方版
	fmt.Printf("LastIndexAny(\"ABCDEABCD\", \"CDE\") = %d\n", strings.LastIndexAny("ABCDEABCD", "CDE"))

	// 指定した文字列で始まる/終わるか検証する
	fmt.Printf("HasPrefix(\"ことほのうみ\", \"こと\") = %t\n", strings.HasPrefix("ことほのうみ", "こと"))
	fmt.Printf("HasPrefix(\"ことほのうみ\", \"ほの\") = %t\n", strings.HasPrefix("ことほのうみ", "ほの"))
	fmt.Printf("HasSuffix(\"ことほのうみ\", \"うみ\") = %t\n", strings.HasSuffix("ことほのうみ", "うみ"))
	fmt.Printf("HasSuffix(\"ことほのうみ\", \"ほの\") = %t\n", strings.HasSuffix("ことほのうみ", "ほの"))
}
