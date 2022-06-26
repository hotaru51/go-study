package foo

import (
	"fmt"
)

type I interface {
	// パッケージ外から参照する場合、インターフェースも同じく先頭が大文字のもののみアクセス可能
	Method1() string
	method2() string
}

type T struct {}

func (t *T) Method1() string {
	return fmt.Sprint("Method1()")
}

func (t *T) method2() string {
	return fmt.Sprint("method2()")
}

type I0 interface {
	Method1() string
}

// インターフェースの入れ子も可能
type I1 interface {
	I0
	Method2() string
}

type I2 interface {
	I1
	Method3() string
}
