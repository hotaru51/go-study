package foo

type Foo struct {
	// 名前の先頭が大文字のものはパッケージの外部から参照可能
	Field1 int
	// こちらは参照不可能
	field2 int
}

// メソッドも同様に名前の先頭が大文字のものはパッケージ外からアクセス可能
func (f *Foo) FooMethod1() int {
	return f.Field1
}

// こちらは不可能
func (f *Foo) fooMethod() int {
	return f.field2
}
