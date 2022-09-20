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

	// 指定した文字列が含まれるか検証
	fmt.Printf("Contains(\"ABCDE\", \"AB\") = %t\n", strings.Contains("ABCDE", "AB"))
	fmt.Printf("Contains(\"ABCDE\", \"XY\") = %t\n", strings.Contains("ABCDE", "XY"))
	fmt.Printf("Contains(\"ABCDE\", \"\") = %t\n", strings.Contains("ABCDE", ""))
	fmt.Printf("Contains(\"\", \"\") = %t\n", strings.Contains("", ""))

	// 指定した文字列のいずれかの文字が含まれているか検証
	fmt.Printf("ContainsAny(\"XYZ\", \"XY\") = %t\n", strings.ContainsAny("XYZ", "XY"))
	fmt.Printf("ContainsAny(\"XYZ\", \"YOSHIKO\") = %t\n", strings.ContainsAny("XYZ", "YOSHIKO"))
	fmt.Printf("ContainsAny(\"XYZ\", \"HANAMARU\") = %t\n", strings.ContainsAny("XYZ", "HANAMARU"))

	// 指定した文字列が何回出現するか検索する
	fmt.Printf("Count(\"ABCABCABCABC\", \"ABC\") = %d\n", strings.Count("ABCABCABCABC", "ABC"))
	fmt.Printf("Count(\"ABCABCABCABC\", \"XYZ\") = %d\n", strings.Count("ABCABCABCABC", "XYZ"))
	fmt.Printf("Count(\"ABC\", \"\") = %d\n", strings.Count("ABC", ""))

	// 文字列を繰り返して結合する
	fmt.Printf("Repeat(\"nico\", 4) = %s\n", strings.Repeat("nico", 4))
	fmt.Printf("Repeat(\"nico\", 0) = %s\n", strings.Repeat("nico", 0))

	// 文字列の置換
	fmt.Printf("Replace(\"AAAAA\", \"A\", \"X\", 1) = %s\n", strings.Replace("AAAAA", "A", "X", 1))
	fmt.Printf("Replace(\"AAAAA\", \"A\", \"X\", 3) = %s\n", strings.Replace("AAAAA", "A", "X", 3))
	fmt.Printf("Replace(\"AAAAA\", \"A\", \"X\", -1) = %s\n", strings.Replace("AAAAA", "A", "X", -1))
	fmt.Printf("Replace(\"JIMOTO\", \"TO\", \"AI\", 1) = %s\n", strings.Replace("JIMOTO", "TO", "AI", 1))

	// split
	fmt.Printf("Split(\"you,chika,ruby,yoshiko\", \",\") = %v\n", strings.Split("you,chika,ruby,yoshiko", ","))
	fmt.Printf("SplitAfter(\"you,chika,ruby,yoshiko\", \",\") = %v\n", strings.SplitAfter("you,chika,ruby,yoshiko", ","))
	fmt.Printf("SplitN(\"you,chika,ruby,yoshiko\", \",\", 3) = %v\n", strings.SplitN("you,chika,ruby,yoshiko", ",", 3))
	fmt.Printf("SplitAfterN(\"you,chika,ruby,yoshiko\", \",\", 3) = %v\n", strings.SplitAfterN("you,chika,ruby,yoshiko", "," ,3))
}
