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
}
