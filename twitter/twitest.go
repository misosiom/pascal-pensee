package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"io"
	"github.com/dghubble/oauth1"
	"github.com/joho/godotenv"
)

func main() {
	// .envファイルを読み込む（1つ上の階層にある場合）
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal(".envの読み込みに失敗しました:", err)
	}

	// 環境変数からAPIキーなどを取得
	apiKey := os.Getenv("API_KEY")
	apiSecret := os.Getenv("API_SECRET")
	accessToken := os.Getenv("ACCESS_TOKEN")
	accessSecret := os.Getenv("ACCESS_SECRET")

	config := oauth1.NewConfig(apiKey, apiSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	url := "https://api.twitter.com/2/tweets"
	body := map[string]string{"text": "パスカルの名言：人間は考える葦である。ww"}
	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatal("リクエスト失敗:", err)
	}
	defer resp.Body.Close()
	defer resp.Body.Close()
	bodyBytes, _ := io.ReadAll(resp.Body)
	fmt.Println("Response body:", string(bodyBytes))
	fmt.Println("ツイート成功：", resp.Status)
}
