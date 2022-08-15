package main

import (
	"fmt"
	"math"
)

func main() {
	// 様々な定数を含む
	fmt.Printf("Pi = %v\n", math.Pi)
	fmt.Printf("Sqrt2 = %v\n", math.Sqrt2)

	// 数値型に関する定数の一部
	fmt.Printf("MaxInt32 = %v\n", math.MaxInt32)
	fmt.Printf("MaxFloat32 = %v\n", math.MaxFloat32)
	fmt.Printf("MaxInt64 = %v\n", math.MaxInt64)
	fmt.Printf("MaxFloat64 = %v\n", math.MaxFloat64)

	// 絶対値を求める
	fmt.Printf("math.Abs(3.14) = %f\n", math.Abs(3.14))
	fmt.Printf("math.Abs(-3.14) = %f\n", math.Abs(-3.14))
}
