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

// PersonとProductにStringifyに定義されたToString()メソッドを実装する
type Stringify interface {
	ToString() string
}

type Person struct {
	Name string
	Age  int
}

func (p *Person) ToString() string {
	return fmt.Sprintf("Name: %s, Age: %d", p.Name, p.Age)
}

type Product struct {
	Name  string
	Price int
}

func (p *Product) ToString() string {
	return fmt.Sprintf("Name: %s, Price: %d", p.Name, p.Price)
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

	person1 := &Person{
		Name: "You Watanabe",
		Age: 16,
	}

	product1 := &Product{
		Name: "Pixel 6a",
		Price: 53900,
	}

	// PersonとProductをStringify配列にまとめてToString()メソッドを実行する
	stringifyArr := [2]Stringify{person1, product1}
	for _, v := range stringifyArr {
		fmt.Println(v.ToString())
	}
}
