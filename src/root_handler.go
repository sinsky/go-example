package hellohandler

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func HandleHello(w http.ResponseWriter, r *http.Request) {
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
}
