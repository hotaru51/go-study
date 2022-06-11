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
}
