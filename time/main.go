package main

import (
	"fmt"
	"time"
)

func main() {
	// 現在時刻の取得
	now := time.Now()
	fmt.Printf("現在時刻: %s", now.String())
}
