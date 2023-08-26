package main

import (
	"net/http"
	echoMessage "sinsky/hello/src"
)

func main() {
	// ハンドラーの設定
	http.HandleFunc("/", echoMessage.HandleHello)

	// サーバーの起動
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
