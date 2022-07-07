package main

import (
	"fmt"
	"path/filepath"
	"os"
)

func main() {
	relativePath := filepath.Dir(os.Args[0])
	absPath, err := filepath.Abs(relativePath)
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println(absPath)
}
