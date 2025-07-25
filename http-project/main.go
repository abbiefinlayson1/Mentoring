package main

import (
	"net/http"
)

// defining a function, takes two arguments - w to write a response and r the http request
func handler(w http.ResponseWriter, r *http.Request) {
	// writes 'ok' back to the client and sends it in the response body
	w.Write([]byte("OK"))
}

func main() {
	// Starts http server on port 3000, uses handler function to handle requests, converts the function to a valid handler.
	http.ListenAndServe(":3000", http.HandlerFunc(handler))
}
