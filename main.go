package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
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
		ids := r.URL.Path[1:] // 例: "1,3,5"
		if ids == "" {
			fmt.Fprintf(w, "No ID provided")
			return
		}
		idStrs := strings.Split(ids, ",")
		var results []string

		for _, idStr := range idStrs {
			id, err := strconv.Atoi(idStr)
			if err != nil {
				results = append(results, fmt.Sprintf("Invalid ID: %s", idStr))
				continue
			}
			if id < 0 || id >= len(passages) {
				results = append(results, fmt.Sprintf("ID out of range: %d", id))
				continue
			}
			results = append(results, fmt.Sprintf("ID %d: %s", id, passages[id].Text))
		}

		for _, res := range results {
			fmt.Fprintln(w, res)
		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
