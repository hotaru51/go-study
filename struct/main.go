package main

import (
	"fmt"
)

type MyInt int

// typeを利用した方のエイリアス
type (
	IntPair     [2]int
	Strings     []string
	AreaMap     map[string][2]float64
	IntsChannel chan []int
)

type Callback func(i int) int

// 構造体の定義
type Point struct {
	X int
	Y int
}

// 型だけの定義も可能
type T struct {
	int
	float64
	string
}

type Person struct {
	name  string
	age   int
}

// 構造体を含む構造体
type SchoolIdle struct {
	person Person
	school string
	grade  int
	group  string
}

// 構造体をフィールド名なしで含む構造体
type Student struct {
	Person
	school string
	grade int
}

func main() {
	var myInt MyInt
	myInt = 1
	fmt.Printf("myInt = %d(%T)\n", myInt, myInt)

	pair := IntPair{1, 2}
	cyaron := Strings{"you", "chika", "ruby"}
	amap := AreaMap{"tokyo": {35.689488, 139.691706}}
	ich := make(IntsChannel, 1)
	ich <- []int{1, 2, 3}
	fmt.Println(<-ich)
	ifArr := [3]interface{}{pair, cyaron, amap}
	for _, v := range ifArr {
		fmt.Println(v)
	}

	sum([3]int{445, 874, 25}, func(res int) int {
		fmt.Printf("callback(): res = %d\n", res)

		return res
	})

	// 構造体型の利用
	var pt Point
	fmt.Printf("pt.X = %d\n", pt.X)
	fmt.Printf("pt.Y = %d\n", pt.Y)
	pt.X = 4
	pt.Y = 5
	fmt.Printf("pt.X = %d\n", pt.X)
	fmt.Printf("pt.Y = %d\n", pt.Y)

	// 定義と同時に値を代入
	pt2 := Point{1, 2}
	pt3 := Point{X: 3, Y: 4}
	fmt.Printf("pt2 = %v, pt3 = %v\n", pt2, pt3)

	t1 := T{
		1,
		3.14,
		"kotori",
	}
	fmt.Printf("T = %v\n", t1)

	// 構造体をフィールド名がついてる場合とついていない場合の比較
	kotori := SchoolIdle{
		person: Person{
			name: "Kotori Minami",
			age: 16,
		},
		school: "Otonokizaka",
		grade: 2,
		group: "μ's",
	}
	fmt.Printf("kotori.person.name = %s\n", kotori.person.name)
	you := Student{
		Person: Person{
			name: "You Watanabe",
			age: 16,
		},
		school: "Uranohoshi",
		grade: 2,
	}
	// フィールド名がない場合はアクセス時のフィールド名は不要
	fmt.Printf("you.name = %s\n", you.name)
}

// callback用の関数型をエイリアスで指定
func sum(intArr [3]int, callback Callback) int {
	res := 0
	for _, v := range intArr {
		res += v
	}

	return callback(res)
}
