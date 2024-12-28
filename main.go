package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This is linkgofer a simple URL shortener"))
	})
	http.ListenAndServe(":8080", nil)
}
