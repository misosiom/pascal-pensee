package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Passage struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

func main() {
	jsonFile, err := os.Open("pensees.json")
	if err != nil {
		fmt.Println("File open error:", err)
		return
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Read error:", err)
		return
	}

	var passages []Passage // ←ここが修正ポイント
	err = json.Unmarshal(byteValue, &passages)
	if err != nil {
		fmt.Println("Unmarshal error:", err)
		return
	}

	if len(passages) == 0 {
		fmt.Println("No passages found.")
		return
	}

	fmt.Println("1件目の文字数:", len(passages[0].Text))
}
