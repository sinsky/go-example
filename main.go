package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Message: "hello",
	}

	// JSONエンコーディング
	data, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// レスポンスヘッダーの設定
	w.Header().Set("Content-Type", "application/json")

	// レスポンスデータの書き込み
	w.Write(data)

	log.Printf("Request received: %s", r.URL.Path)
}

func main() {
	// ハンドラーの設定
	http.HandleFunc("/", helloHandler)

	// サーバーの起動
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
