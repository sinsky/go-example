package router

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func HandleHello(w http.ResponseWriter, r *http.Request) {
	// HandleHello handles HTTP requests and returns a JSON response with a "hello" message.
	//
	// Parameters:
	// - w http.ResponseWriter: The response writer used to write the response data.
	// - r *http.Request: The HTTP request received from the client.
	//
	// Returns:
	// - None. The function writes the JSON data to the response writer.

	response := Response{
		Message: "hello",
	}

	// JSON encoding
	data, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the "Content-Type" header of the response to "application/json".
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON data to the response writer.
	w.Write(data)
}
