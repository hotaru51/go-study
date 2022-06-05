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

	// 値渡しと参照渡し、配列は値渡しでスライスは参照渡し
	arr2 := [3]int{1, 2, 3}
	powArray(arr2)
	fmt.Printf("arr2 = %v\n", arr2)
	s10 := []int{1, 2, 3}
	powSlice(s10)
	fmt.Printf("s10 = %v\n", s10)

	// 配列とスライスで初期化をしない場合の初期値の違い
	var (
		arr3 [3]int
		s11 []int
	)
	fmt.Printf("arr3 = %v, s11 = %v\n", arr3, s11)
	fmt.Printf("s11 == nil -> %v\n", s11 == nil)

	// map
	// var m map[int]string
	// makeで生成する場合
	m := make(map[int]string)
	m[1] = "you"
	m[2] = "chika"
	m[3] = "ruby"
	fmt.Println(m)

	// 宣言と同時に初期化
	m2 := map[string]string{
		"yoshiko": "Yoshiko Tsushima",
		"riko": "Riko Sakurauchi",
		"mari": "Mari Ohara", // <- 複数行で書くときに最後のカンマが必須なので注意
	}
	fmt.Println(m2)

	// 存在しないキーの要素を参照する場合の挙動
	m3 := map[string]string{"kashiyuka": "Yuka Kashino"}
	fmt.Printf("m3 = %v\n", m3)
	// この場合、このキーの要素は存在しないため、string型の初期値の空文字が返る
	ayaka := m3["ayaka"]
	fmt.Printf("ayaka = %s\n", ayaka)
	// 下記のように参照することでキーが存在するか判定可能
	nocchi, ok := m3["nocchi"]
	fmt.Printf("nocchi = %s(%T), ok = %v(%T)\n", nocchi, nocchi, ok, ok)
}

// 可変長引数はスライスとして受け取る
func variableLengthArgumentTest(s ...string) {
	fmt.Printf("s = %v(%T)\n", s, s)
}

func powArray(a [3]int) {
	for i, v := range a {
		a[i] = v * v
	}
	fmt.Printf("powArray(a): %v\n", a)
}

func powSlice(s []int) {
	for i, v := range s {
		s[i] = v * v
	}
	fmt.Printf("powSlice(a): %v\n", s)
}
