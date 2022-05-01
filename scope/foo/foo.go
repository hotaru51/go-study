package foo

const (
    // パッケージ外から参照できる
    A = 445
    // パッケージ外から参照できない
    b = 745
)

// パッケージ外から参照できる
func TestFunction1() string {
    return "nico"
}

// パッケージ外から参照できない
func testFunction2() string {
    return "maki"
}
