package main

import (
	"fmt"
	"os"
	// "log"
)

func openFile(path string) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	// deferで開いたファイルを確実にcloseする
	fmt.Println(f.Name())
	defer f.Close()
}

func main() {
	// os.Exitで終了するとdeferは実行されない
	defer func() {
		fmt.Println("defer")
	}()

	// 存在しないファイルを開こうとした場合にエラーメッセージを出力
	_, err := os.Open("hoge")
	if err != nil {
		// log.Fatal(err)
	}

	// コマンドライン引数はos.Args[]で参照できる
	for i, v := range os.Args {
		fmt.Printf("os.Args[%d] = \"%s\"\n", i, v)
	}

	openFile("./uranohoshi.txt")

	// 環境変数を参照する
	fmt.Println(os.Getenv("PATH"))
	// 存在しない場合は空文字
	fmt.Printf("%#v\n", os.Getenv("HOGE"))

	// 厳密に環境変数が存在するかチェックする場合はos.LookupEnv()を使用する
	if testEnv, ok := os.LookupEnv("test_env"); ok {
		fmt.Println(testEnv)
	} else {
		fmt.Println("no $test_env")
	}

	// 実行ファイルの絶対パス取得
	executablePath, _ := os.Executable()
	fmt.Println(executablePath)

	// 新規ファイルの作成
	newf, _ := os.Create("otonokizaka.txt")
	fs, _ := newf.Stat()
	fmt.Println(fs.Name())
	fmt.Println(fs.Size())
	fmt.Println(fs.IsDir())
	// ファイルに[]byte型を書き込み
	newf.Write([]byte("Honoka Kousaka\n"))

	// パーミッション等を指定してファイルオープン
	os.OpenFile("./nijigasaki.txt", os.O_CREATE, 0600)

	// ファイルの削除
	err = os.Remove("./kasukasu.txt")
	if err != nil {
		fmt.Println(err)
	}

	// ファイルを含むディレクトリを削除する場合はos.RemoveAllを使用する
	err = os.RemoveAll("./yohane_dir")
	if err != nil {
		fmt.Println(err)
	}

	// ファイル名の変更
	os.Create("./yohane.txt")
	os.Rename("./yohane.txt", "./yoshiko.txt")

	// カレントディレクトリの取得
	dir, err := os.Getwd(); if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dir)

	// ディレクトリを開く
	d1, err := os.Open("."); if err != nil {
		fmt.Println(err)
	}

	// 開いたディレクトリの内容を取得
	fdir, err := d1.ReadDir(0)

	// ディレクトリ内のファイル名をすべて表示する
	for _, f := range fdir {
		fmt.Println(f.Name())
	}

	// システムのテンポラリディレクトリの取得
	fmt.Println(os.TempDir())

	// シンボリックリンクの作成
	err = os.Symlink("./uranohoshi.txt", "./aqours.txt")

	// シンボリックリンク先のファイルのパスを読み込む
	sp, err := os.Readlink("./aqours.txt"); if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sp)

	os.Exit(0)
}
