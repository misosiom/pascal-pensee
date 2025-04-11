package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dghubble/oauth1"
	"github.com/joho/godotenv"
)

func main() {
	// .env 読み込み
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal(".envの読み込みに失敗:", err)
	}

	apiKey := os.Getenv("API_KEY")
	apiSecret := os.Getenv("API_SECRET")
	accessToken := os.Getenv("ACCESS_TOKEN")
	accessSecret := os.Getenv("ACCESS_SECRET")

	config := oauth1.NewConfig(apiKey, apiSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	url := "https://api.twitter.com/2/tweets"

	// 最初の投稿
	tweetID := postTweet(httpClient, url, "パスカルの名言：人間は考える葦である。", "")

	// 返信チェーン（前の返信IDに次の返信を繋げる）
	for i := 1; i <= 5; i++ {
		replyText := fmt.Sprintf("これは返信%dです。", i)
		tweetID = postTweet(httpClient, url, replyText, tweetID)
		time.Sleep(1 * time.Second) // 過剰投稿防止
	}
}

func postTweet(client *http.Client, url string, text string, replyToID string) string {
	// リクエストボディ作成
	body := map[string]interface{}{
		"text": text,
	}
	if replyToID != "" {
		body["reply"] = map[string]string{
			"in_reply_to_tweet_id": replyToID,
		}
	}

	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("ツイート失敗:", err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	fmt.Println("レスポンス:", string(respBody))

	var res struct {
		Data struct {
			ID string `json:"id"`
		} `json:"data"`
	}
	json.Unmarshal(respBody, &res)

	return res.Data.ID
}
