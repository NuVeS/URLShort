package main

import (
	"net/http"

	"github.com/NuVeS/URLShort/cmd/router"
)

func main() {
	router := router.NewRouter()

	http.ListenAndServe(":8080", router)

	// print(shortener.MakeShort("test"))

}
