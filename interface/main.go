package main

import (
	"fmt"
)

type MyError struct {
	Message string
	ErrCode int
}

func (e *MyError) Error() string {
	return e.Message
}

func RaiseError() error {
	return &MyError{Message: "エラーが発生しました", ErrCode: 745}
}

func main() {
	// 本来のデータ型はMyErrorだけど、errはerror型となる
	err := RaiseError()
	fmt.Println(err.Error())
	// error型なので、この段階ではErrCodeは取り出せない
	// fmt.Println(e.ErrCode)

	// 型アサーションでMyErrorを取り出すことでErrCodeもアクセス可能
	e, ok := err.(*MyError)
	if ok {
		fmt.Println(e.ErrCode)
	}
}
