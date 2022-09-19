package main

import (
	"net/http"

	"github.com/NuVeS/URLShort/cmd/handler"
)

func main() {
	http.HandleFunc("/", handler.MainHandler)

	err := http.ListenAndServe(":33333", nil)

	if err != nil {
		print("Error starting server")
	}

	// print(shortener.MakeShort("test"))
}
