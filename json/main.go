package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"path/filepath"
	"log"
)

func readJsonFile(path string) string {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	jsonText, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	return string(jsonText)
}

func main() {
	executablePath, _ := os.Executable()
	jsonPath := filepath.Dir(executablePath) + "/test.json"
	jsonText := readJsonFile(jsonPath)
	fmt.Println(jsonText)
}
