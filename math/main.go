package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
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

	// 累乗を求める
	fmt.Printf("math.Pow(2, 8) = %f\n", math.Pow(2, 8))

	// 平方根
	fmt.Printf("math.Sqrt(4) = %f\n", math.Sqrt(4))

	// 最大値と最小値
	ysk := 445
	nsk := 745
	fmt.Printf("math.Max(ysk, nsk) = %f\n", math.Max(float64(ysk), float64(nsk)))
	fmt.Printf("math.Min(ysk, nsk) = %f\n", math.Min(float64(ysk), float64(nsk)))

	// 小数点の切り捨て
	fmt.Printf("math.Trunc(1.5) = %f\n", math.Trunc(1.5))
	fmt.Printf("math.Trunc(-1.5) = %f\n", math.Trunc(-1.5))

	// 与えた数値より小さい整数を返す
	fmt.Printf("math.Floor(1.5) = %f\n", math.Floor(1.5))
	fmt.Printf("math.Floor(-1.5) = %f\n", math.Floor(-1.5))

	// 与えた数値より大きい整数を返す
	fmt.Printf("math.Ceil(1.5) = %f\n", math.Ceil(1.5))
	fmt.Printf("math.Ceil(-1.5) = %f\n", math.Ceil(-1.5))

	// 非数と非数の確認
	fmt.Printf("math.Sqrt(-1) = %v\n", math.Sqrt(-1))
	fmt.Printf("math.IsNaN(math.Sqrt(-1)) = %v\n", math.IsNaN(math.Sqrt(-1)))

	// 無限大と負の無限大
	fmt.Printf("math.Inf(0) = %v\n", math.Inf(0))
	fmt.Printf("math.Inf(-1) = %v\n", math.Inf(-1))

	// 乱数の生成
	rand.Seed(256)
	fmt.Printf("rand.Float64() = %f\n", rand.Float64())
	// シードの設定に現在時刻を使う
	rand.Seed(time.Now().UnixNano())
	// 0〜99の乱数を返す
	fmt.Printf("rand.Intn(100) = %d\n", rand.Intn(100))

	// 擬似乱数生成器の生成
	// 疑似乱数生成器のソースを生成
	rndSrc := rand.NewSource(time.Now().UnixNano())
	// ソースを元に疑似乱数生成器を生成
	myRnd := rand.New(rndSrc)
	// 生成した擬似乱数生成器を使用して乱数を生成
	fmt.Printf("myRnd.Intn(100) = %d\n", myRnd.Intn(100))
}
