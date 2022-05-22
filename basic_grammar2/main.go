package main

import (
	"fmt"
)

func main() {
	// for
	// 下記は無限ループ
	/*
	   for {
	       fmt.Println("逮捕")
	   }
	*/

	// いつもの
	member1 := [...]string{"chika", "riko", "you"}
	for i := 0; i < 3; i++ {
		fmt.Println(member1[i])
	}

	// if
	m1 := 1
	// m1 := 2
	// m1 := 3
	if m1 == 1 {
		fmt.Println("m1 == 1")
	} else if m1 == 2 {
		fmt.Println("m1 == 2")
	} else {
		fmt.Println("m1 != 1 && m1 != 2")
	}

	// 簡易文付きif
	// ここで定義した変数はブロック内のみ有効
	if x, y := 1, 2; x < y {
		fmt.Printf("%d < %d\n", x, y)
	}

	// 他の言語のwhileの様なfor
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	// rangeを使用したfor
	for i, s := range member1 {
		fmt.Printf("member1[%d] = %s\n", i, s)
	}

	// stringに対してrangeを使用するとrune型が返る
	name := "Kotori Minami"
	fmt.Printf("name = %s(%T)\n", name, name)
	for i, r := range name {
		fmt.Printf("name[%d] = %v(%T) = %s\n", i, r, r, string(r))
	}

	// switch
	sw1 := 3
	switch sw1 {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("unknown")
	}

	// CやJavaのswitchのようにフォールスルーさせる
	sw2 := "N"
	switch sw2 {
	case "N":
		sw2 += "I"
		fallthrough
	case "NI":
		sw2 += "C"
		fallthrough
	default:
		sw2 += "O"
	}
	fmt.Println(sw2)

	// 簡易分付きswitch
	// if同様、下記の変数はswitch内でのみ有効
	switch member := "you"; member {
	case "yoshiko", "ruby", "hanamaru":
		fmt.Printf("%s = 1年生\n", member)
	case "chika", "you", "riko":
		fmt.Printf("%s = 2年生\n", member)
	case "kanan", "mari", "dia":
		fmt.Printf("%s = 3年生\n", member)
	default:
		fmt.Println("unknown")
	}

	// 型アサーション
	var if1 interface{} = 3
	chk1 := if1.(int)     // chk1にはint型で3が入る
	// chk2 := if1.(float64) // これは実行時エラー
	fmt.Printf("%d(%T)\n", chk1, chk1)
	// fmt.Println(chk1, chk2)

	// 型アサーションのエラー検知
	var if2 interface{} = 3
	chk2, isInt := if2.(int)
	chk3, isFloat64 := if2.(float64)
	// int型のアサーションが成功、isIntにはtrueが入る
	fmt.Printf("if2.(int) = %v, %v\n", chk2, isInt)
	// float64型のアサーションに失敗、isFloat64にはfalseが入る
	fmt.Printf("if2.(float) = %v, %v\n", chk3, isFloat64)

	// 型によるswitch
	var if3 interface{} = "ﾊﾉｹﾁｪﾝ"
	switch if3.(type) {
	case bool:
		fmt.Println("if3 = boll")
	case int:
		fmt.Println("if3 = int")
	case string:
		fmt.Println("if3 = string")
	default:
		fmt.Println("unknown")
	}

	// ラベル付き文
	// 深いループから抜ける
	LOOP:
		for {
			fmt.Println("loop1")
			for {
				fmt.Println("loop2")
				for {
					fmt.Println("loop3")
					break LOOP
				}
			}
		}

	L:
		for i := 1; i <= 3; i++ {
			for j := 1; j<= 3; j++ {
				if j > 1 {
					continue L
				}
				fmt.Printf("%d * %d = %d\n", i, j, i *j)
			}
			fmt.Println("yohane")
		}

	// defer
	runDefer()
}

func runDefer() {
	// defer文の式は関数終了後に処理される
	defer fmt.Println("defer")
	fmt.Println("mikan")
}
