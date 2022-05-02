package main

import (
    "fmt"
    "./foo"
    // エイリアスをつけて参照することも可能
    f "./foo"
    // パッケージ名なしで参照できるようにインポートすることも可能
    // ただし、インポートしたパッケージと識別子が衝突しないように注意
    . "./foo"
)

func main() {
    fmt.Printf("foo.A = %d\n", int(foo.A))
    // これは参照できない
    // fmt.Printf("foo.b = %d\n", int(foo.b))

    fmt.Printf("TestFunction1() => %s\n", foo.TestFunction1())
    // これは参照できない
    // fmt.Printf("testFunction2() => %s", foo.testFunction2())
    fmt.Println(f.TestFunction1())
    fmt.Println(TestFunction1())

    fmt.Println(functionScopeSample(1))
}

func functionScopeSample(a int) (b int) {
    // これは当然エラー
    // a := 1
    c := 1

    // 戻り値の変数も定義済みであればエラー
    // var b

    {
        // {}で独立したブロックを定義することで重複した変数を使うことができる
        // ただし、有効範囲はこのブロック内のみ
        var b string
        b = "kashiyuka"
        fmt.Println(b)
    }

    // 以下のbは戻り値のint型のb
    fmt.Println(b)

    return c
}
