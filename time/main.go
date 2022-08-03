package main

import (
	"fmt"
	"log"
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

	// 時刻の間隔の表現
	{
		// 4時間45分を文字列から生成
		d, _ := time.ParseDuration("4h45m")
		fmt.Println(d)

		// 現在から4時間45分後を生成
		t := time.Now()
		fmt.Printf("現在時刻        : %v\n", t)
		t = t.Add(d)
		fmt.Printf("現在時刻 + 4h45m: %v\n", t)
	}

	// 時刻の差分の取得
	{
		ybd := time.Date(2022, 4, 17, 0, 0, 0, 0, time.Local)
		now := time.Now()
		fmt.Println(ybd.Sub(now))
	}

	// 日付を比較する
	{
		t1 := time.Date(2022, 4, 17, 0, 0, 0, 0, time.Local)
		t2 := time.Date(2022, 8, 1, 0, 0, 0, 0, time.Local)
		fmt.Printf("t1 = %#v\n", t1)
		fmt.Printf("t2 = %#v\n", t2)
		fmt.Printf("t1 < t2 => %#v\n", t1.Before(t2))
		fmt.Printf("t1 > t2 => %#v\n", t1.After(t2))
		fmt.Printf("t2 < t1 => %#v\n", t2.Before(t1))
		fmt.Printf("t2 > t1 => %#v\n", t2.After(t1))
	}

	// 日付の増減
	{
		t := time.Date(2022, 3, 20, 0, 0, 0, 0, time.Local)
		fmt.Println(t)
		t = t.AddDate(0, 1, -3)
		fmt.Println(t)
	}

	// 文字列から時刻を生成する
	{
		// 第一引数がフォーマット代わり
		if t, err := time.Parse("2006/01/02", "2022/08/01"); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(t)
		}

		// フォーマットは定数でも用意されている
		if t, err := time.Parse(time.RFC3339, "2022-04-17T00:00:00+09:00"); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(t)
		}
	}
}
