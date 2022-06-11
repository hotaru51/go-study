package main

import (
	"fmt"
)

func main() {
	// ポインタ型の初期値はnil
	var p *int
	fmt.Printf("(p == nil) = %v\n", (p == nil))

	// 変数のアドレス取得
	i := 5
	pi := &i
	fmt.Printf("pi => %T\n", pi)

	// 関数への参照渡し
	i2 := 1
	fmt.Printf("i2 = %d\n", i2)
	inc(&i2)
	fmt.Printf("i2 = %d\n", i2)
	inc(&i2)
	fmt.Printf("i2 = %d\n", i2)
	inc(&i2)
	fmt.Printf("i2 = %d\n", i2)

	// 配列の参照
	intArr := [3]int{1, 2, 3}
	fmt.Printf("intArr = %v\n", intArr)
	pow(&intArr)
	fmt.Printf("intArr = %v\n", intArr)

	// 配列参照時の記載の省略
	strArr := [3]string{"you", "chika", "ruby"}
	pstrArr := &strArr
	fmt.Println(strArr[0])
	// (*pstrArr)[0] という記載でなくても自動で参照してくれる
	fmt.Println(pstrArr[0])
	fmt.Println(pstrArr[1])
	fmt.Println(pstrArr[2])

	// rangeでも省略可能
	for _, str := range pstrArr {
		fmt.Println(str)
	}
	// len、cap等も省略可能
	fmt.Printf("len(pstrArr) = %v, cap(pstrArr) = %v, pstrArr[0:2] = %v\n",
		len(pstrArr), cap(pstrArr), pstrArr[0:2],
	)
}

func inc(i *int) {
	*i++
}

func pow(p *[3]int) {
	i := 0
	for i < 3 {
		(*p)[i] = (*p)[i] * (*p)[i]
		i++
	}
}
