package urls

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/NuVeS/URLShort/cmd/models"
	"github.com/NuVeS/URLShort/cmd/shortener"
	"github.com/NuVeS/URLShort/cmd/storage"
)

var DB storage.StorageAPI

var addingPart = "http://localhost/" // os.Getenv("REMOTE_ADDR")

func Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	token := r.Header.Get("token")

	if DB.GetUserByToken(token) == nil {
		sendError(w, http.StatusUnauthorized)
		return
	}

	var body models.LinkRequest
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		sendError(w, http.StatusInternalServerError)
		return
	}

	ok := DB.DeleteLink(token, body.Url)

	if ok {
		sendOK(w)
	} else {
		sendError(w, http.StatusOK)
	}

}

func List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	token := r.Header.Get("token")

	if DB.GetUserByToken(token) == nil {
		sendError(w, http.StatusUnauthorized)
		return
	}

	list := DB.ListLinks(token)
	var jsonList = make([]models.LinkResponse, 0)
	for _, item := range list {
		temp := models.LinkResponse{Url: item.Url, ShortUrl: addingPart + item.Short}
		jsonList = append(jsonList, temp)
	}
	response := models.LinkListResponse{List: jsonList}
	json, err := json.Marshal(response)

	if err == nil {
		w.WriteHeader(http.StatusOK)
		w.Write(json)
	} else {
		sendError(w, http.StatusInternalServerError)
	}
}

func Shorten(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	token := r.Header.Get("token")

	if DB.GetUserByToken(token) == nil {
		sendError(w, http.StatusUnauthorized)
		return
	}

	var body models.NewLinkRequest
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		sendError(w, http.StatusInternalServerError)
		return
	}

	writeLinkResponse := func(beauty *string) {
		link := makeNewLink(token, body.Url, beauty)
		response := models.LinkResponse{Url: body.Url, ShortUrl: addingPart + link}
		json, _ := json.Marshal(response)

		w.WriteHeader(http.StatusOK)
		w.Write(json)
	}

	if len(body.Beauty) != 0 {
		available := DB.IsAvailable(body.Beauty)
		if !available {
			w.WriteHeader(http.StatusBadRequest)
			json, _ := json.Marshal("Not available link")
			w.Write(json)
		} else {
			writeLinkResponse(&body.Beauty)
		}
	} else {
		writeLinkResponse(nil)
	}

}

func Route(w http.ResponseWriter, r *http.Request) {
	uri := r.RequestURI
	paths := strings.Split(uri, "/")
	url := DB.GetURL(paths[len(paths)-1])

	http.Redirect(w, r, url, http.StatusSeeOther)
}

func makeNewLink(token string, link string, beauty *string) string {
	if beauty != nil {
		DB.AddLink(token, link, *beauty)
		return *beauty
	} else {
		short := shortener.MakeShort(link)
		DB.AddLink(token, link, short)
		return short
	}
}

func sendError(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
	json, err := json.Marshal("Failed")
	if err == nil {
		w.Write(json)
	}
}

func sendOK(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	json, _ := json.Marshal("OK")
	w.Write(json)
}
