package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"path/filepath"
	"log"
	"encoding/json"
)

// privateなフィールドにしてしまうとUnmarshal時に値が入らなくなるので注意
type YouChika struct {
	You   string
	Chika string
}

func (yc *YouChika) String() string {
	return fmt.Sprintf("you: %s, chika: %s\n", yc.You, yc.Chika)
}

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
	var yc YouChika
	err := json.Unmarshal([]byte(jsonText), &yc)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(&yc)
}
