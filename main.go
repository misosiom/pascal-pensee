package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"net/http"
	"log"
	"strconv"
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

	fmt.Println("全体の長さ:", len(passages))
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		id, error := strconv.Atoi(r.URL.Path[1:])
		if error != nil {
			fmt.Fprintf(w, "Invalid ID format")
			return
		}
		if id < 0 || id >= len(passages) {
			fmt.Fprintf(w, "ID out of range")
			return
		}
		fmt.Fprintf(w, "The passage which have the id is: %s\n", passages[id].Text)
	})
    log.Fatal(http.ListenAndServe(":8080", nil))
}
