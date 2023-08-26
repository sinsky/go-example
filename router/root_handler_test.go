package router

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// To verify that the HandleHello function returns a JSON response with a "hello" message.
func TestHandleHelloReturnsHelloMessage(t *testing.T) {
	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP response recorder
	rr := httptest.NewRecorder()

	// Call the HandleHello function with the response recorder and request
	HandleHello(rr, req)

	// Check the response status code
	if rr.Code != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, rr.Code)
	}

	// Check the response body
	expected := `{"message":"hello"}`
	if rr.Body.String() != expected {
		t.Errorf("expected body %s but got %s", expected, rr.Body.String())
	}
}

// To verify that the "Content-Type" header of the response is set to "application/json".
func TestSetContentTypeHeader(t *testing.T) {
	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP response recorder
	rr := httptest.NewRecorder()

	// Call the HandleHello function with the response recorder and request
	HandleHello(rr, req)

	// Check if the "Content-Type" header is set to "application/json"
	if rr.Header().Get("Content-Type") != "application/json" {
		t.Errorf("Expected Content-Type header to be 'application/json', got '%s'", rr.Header().Get("Content-Type"))
	}
}
