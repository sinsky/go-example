package main

import (
	"hello/module/router"
	"net/http"
)

func main() {
	// ハンドラーの設定
	http.HandleFunc("/", router.HandleHello)

	// サーバーの起動
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
