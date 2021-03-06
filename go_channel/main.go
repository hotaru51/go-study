package main

import (
	"fmt"
)

func main() {
	// チャンネルの定義
	var ch chan int
	// 受信専用チャンネル
	// var receiveOnly <-chan int
	// 送信専用チャンネル
	//var sendOnly chan<- int

	// makeでの生成
	ch = make(chan int)
	// サイズ指定も可能
	ch = make(chan int, 8)
	// 送信用、受信用
	// receiveOnly = make(<-chan int, 5)
	// sendOnly = make(chan<- int, 5)

	// チャンネルへの値送信
	ch <- 5
	// チャンネルから値取得
	i := <-ch
	fmt.Printf("i = %d\n", i)

	// この状態で下記のコメントアウトを外すと永久に受信待ちでランタイムエラーとなる
	// fmt.Println(<-ch)

	// goroutineで受信
	go receiver(ch)

	// バッファのサイズはcapで参照できる
	fmt.Printf("cap(ch) = %d\n", cap(ch))
	for i := 0; i < 10; i++ {
		ch <- i
		// バッファ内に格納されている要素数はlenで取れる
		fmt.Printf("len(ch) = %d\n", len(ch))
	}

	close(ch)
	// 下記のコメントを外してcloseしたチャンネルに対してデータを送信するとランタイムエラーとなる
	// ch <- 5

	// closeされたチャンネルが空であることを確認する
	closeTestCh := make(chan int)
	close(closeTestCh)
	i, ok := <-closeTestCh
	fmt.Printf("i = %d, ok = %v\n", i, ok)

	// チャンネルとfor
	// 下記でチャンネルから値を取得できるが、クローズを検知できない
	/*
	forTestCh := make(chan int)
	forTestCh <- 1
	forTestCh <- 2
	forTestCh <- 3
	forTestCh <- 4
	for i := range forTestCh {
		fmt.Println(i)
	}
	*/

	// select
	ch1 := make(chan int, 5)
	ch2 := make(chan int, 5)

	go func() {
		for {
			select {
			case e1 := <-ch1:
				fmt.Println(e1)
			case e2 := <-ch2:
				fmt.Println(e2)
			}
		}
	}()

	for i := 0; i < 10; i++ {
		ch1 <- 11
		ch2 <- 22
	}
}

func receiver(ch <-chan int) {
	for {
		i := <-ch
		fmt.Println(i)
	}
}
