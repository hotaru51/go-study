package main

import (
	"fmt"
	"time"
)

func main() {
	// 現在時刻の取得
	now := time.Now()
	fmt.Printf("現在時刻: %s", now.String())

	// 指定日付の生成と表示
	{
		t := time.Date(2022, 4, 17, 10, 18, 0, 0, time.Local)
		fmt.Println(t.Year())
		// time.Month()はtime.Month型が返されるが、実体はint型
		fmt.Println(t.Month())
		// 下記のようにすると4が返される
		// fmt.Println(int(t.Month()))
		fmt.Println(t.Day())
		// 曜日も月と同様、time.Weekday型が存在する
		fmt.Println(t.Weekday())
		// 日曜日なので0が返される
		// fmt.Println(int(t.Weekday()))
		fmt.Println(t.Hour())
		fmt.Println(t.Minute())
		fmt.Println(t.Second())
	}
}
