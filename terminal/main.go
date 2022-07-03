package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("os.Stdin.Fd() = %v\n", os.Stdin.Fd())
	for scanner.Scan() {
		fmt.Println(string(scanner.Bytes()))
	}
}
