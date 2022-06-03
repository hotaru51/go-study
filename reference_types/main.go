package main

import (
	"fmt"
)

func main() {
	//スライス
	var s []int
	// makeで生成
	s = make([]int, 10)
	fmt.Println(s)
	// 要素数を調べる
	fmt.Printf("len(s) = %d\n", len(s))

	// 要素数5、容量10のすらいすを生成
	s1 := make([]int, 5, 10)
	fmt.Printf("len(s1) = %d, cap(s1) = %d\n", len(s1), cap(s1))

	// makeを使用せずにスライス生成
	s2 := []string{"you", "chika", "riko"}
	fmt.Printf("s2 = %v\n", s2)
	arr1 := [5]string{"honoka", "kotori", "nico", "maki", "umi"}
	s3 := arr1[2:4]
	fmt.Printf("arr1 = %v\n", arr1)
	fmt.Printf("s3 = %v\n", s3)

	// スライスに要素を追加
	s4 := []string{"ruby", "hanamaru"}
	fmt.Printf("s4 = %v\n", s4)
	s4 = append(s4, "yoshiko")
	fmt.Printf("s4 = %v\n", s4)
	// スライスにスライスを追加することも可能
	s5 := []string{"kanan", "mari", "dia"}
	s6 := append(s2, append(s4, s5...)...)
	fmt.Printf("s6 = %v\n", s6)

	// スライスのコピー
	// コピー先の容量を超えるものは切り捨てられる
	s7 := copy(s4, s6)
	fmt.Printf("s4 = %v, s6 = %v, s7 = %v\n", s4, s6, s7)

	intSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// 簡易スライス式
	s8 := intSlice[2:4]
	fmt.Printf("s8 = %v\n", s8)
	fmt.Printf("len(s8) = %d, cap(s8) = %d\n", len(s8), cap(s8))
	// 完全スライス式
	s9 := intSlice[2:4:4]
	fmt.Printf("s9 = %v\n", s9)
	fmt.Printf("len(s9) = %d, cap(s9) = %d\n", len(s9), cap(s9))

	// forでスライスを回す
	for i, s := range s6 {
		fmt.Printf("s6[%d] = %s\n", i, s)
		// rangeの場合はここで要素を追加しても無限ループにはならない
		s6 = append(s6, "yohane")
	}
	fmt.Printf("s6 = %v\n", s6)

	variableLengthArgumentTest("honoka", "kotori", "umi")

	// スライスを可変長引数として渡すことも可能
	variableLengthArgumentTest([]string{"nico", "maki", "eli"}...)
}

// 可変長引数はスライスとして受け取る
func variableLengthArgumentTest(s ...string) {
	fmt.Printf("s = %v(%T)\n", s, s)
}