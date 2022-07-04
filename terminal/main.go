package main

import (
	"fmt"
	"os"
	"io/ioutil"

	"golang.org/x/term"
)

func main() {
	if term.IsTerminal(0) {
		fmt.Println("terminal")
	} else {
		fmt.Println("pipe")
		data, err := ioutil.ReadAll(os.Stdin); if err != nil {
			fmt.Println(err)
		}
		fmt.Print(string(data))
	}
}
