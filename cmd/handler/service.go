package handler

import (
	"encoding/json"
	"github.com/NuVeS/URLShort/cmd/shortener"
	"net/http"
)

var urls = make(map[string]string, 0)

func MainHandler(writer http.ResponseWriter, request *http.Request) {
	print("Got request")
	if request.Method == http.MethodPost {
		post(writer, request)
	} else if request.Method == http.MethodGet {
		get(writer, request)
	}
}

func post(writer http.ResponseWriter, request *http.Request) {
	var body urlRequest
	err := json.NewDecoder(request.Body).Decode(&body)
	if err != nil {
		sendError(writer)
		return
	}
	short := shortener.MakeShort(body.Url)
	urls[short] = body.Url
	sendResponse(short, writer)
}

func get(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	if !query.Has("url") {
		sendError(writer)
		return
	}
	url := query.Get("url")
	sendResponse(urls[url], writer)
}

func sendError(writer http.ResponseWriter) {
	var error = errorMessage{"Failed"}
	json, err := json.Marshal(error)
	if err == nil {
		writer.Write(json)
	}
}

func sendResponse(response string, writer http.ResponseWriter) {
	var message = urlRequest{response}
	json, err := json.Marshal(message)
	if err == nil {
		writer.Write(json)
	}
}
