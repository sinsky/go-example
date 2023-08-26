package main

import (
	"net/http"
	"sinsky/hello/router"
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
