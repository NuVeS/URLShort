package auth

import (
	"encoding/json"
	"net/http"

	"github.com/NuVeS/URLShort/cmd/models"
	"github.com/NuVeS/URLShort/cmd/storage"
	"github.com/google/uuid"
)

var DB storage.StorageAPI

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var body models.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		sendError(w)
		return
	}

	ok, user := DB.GetUser(body.Name)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		sendError(w)
		return
	}

	token := uuid.New().String()
	DB.SetToken(token, user)

	response := models.LoginResponse{Token: token}

	json, err := json.Marshal(response)
	if err == nil {
		w.WriteHeader(http.StatusOK)
		w.Write(json)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		sendError(w)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	token := r.Header.Get("token")

	user := DB.GetUserByToken(token)
	if user == nil {
		w.WriteHeader(http.StatusInternalServerError)
		sendError(w)
		return
	}

	ok := DB.DeleteUser(user)

	if ok {
		w.WriteHeader(http.StatusOK)
		message, _ := json.Marshal("OK")
		w.Write(message)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		sendError(w)
	}

}

func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var body models.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		sendError(w)
		return
	}

	ok := DB.CreateUser(body.Name, body.Password)

	if ok {
		token := uuid.New().String()
		response := models.LoginResponse{Token: token}
		json, _ := json.Marshal(response)

		w.WriteHeader(http.StatusOK)
		w.Write(json)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		sendError(w)
	}
}

func sendError(writer http.ResponseWriter) {
	json, err := json.Marshal("Failed")
	if err == nil {
		writer.Write(json)
	}
}
