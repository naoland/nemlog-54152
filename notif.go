//usr/bin/env go run $0 $@ ; exit

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// How to use
//
// > go run notif.go -token $LINE_TOKEN message
func main() {
	token := flag.String("token", "", "LINE Notify Personal Access Token")
	flag.Parse()

	if *token == "" {
		log.Fatal("not token")
	}

	message := flag.Arg(0)
	if message == "" {
		log.Fatal("not text")
	}

	fmt.Println("token:", token)
	fmt.Println("message:", message)

	results, err := sendMessage(*token, message)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("%s\n", results)
	// {"status":200,"message":"ok"}
}

// sendMessage LINE通知にメッセージを送信します。
//
// token パーソナルアクセストークン
// message 送信するメッセージ
func sendMessage(token, text string) (results []byte, err error) {
	data := url.Values{"message": {text}}
	api := "https://notify-api.line.me/api/notify"

	req, err := http.NewRequest("POST", api, strings.NewReader(data.Encode()))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	results, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return
}
