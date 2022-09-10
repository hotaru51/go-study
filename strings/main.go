package main

import (
	"fmt"
	"strings"
)

func main() {
	// []string内の要素を結合する
	cyaron := []string{"you", "chika", "ruby"}
	fmt.Println(strings.Join(cyaron, "! ") + "!")
}
