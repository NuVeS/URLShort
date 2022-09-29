package auth

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/NuVeS/URLShort/cmd/models"
	"github.com/NuVeS/URLShort/cmd/storage"
)

var lastToken string

type routeFunc = func(w http.ResponseWriter, r *http.Request)

func init() {
	DB = &storage.Storage
}

func TestRegister(t *testing.T) {
	rr := makeTest(t,
		"POST",
		"/register",
		"",
		"{\"name\": \"root\",\"password\": \"123\"}",
		Register,
	)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	var body models.LoginResponse
	_ = json.NewDecoder(rr.Body).Decode(&body)
	if token := body.Token; len(token) == 0 {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), token)
	} else {
		lastToken = token
	}
}

func TestErrRegister(t *testing.T) {
	_ = DB.CreateUser("root", "123")
	rr := makeTest(t,
		"POST",
		"/register",
		"",
		"{\"name\": \"root\",\"password\": \"123\"}",
		Register,
	)

	// Check the status code is what we expect.
	if status := rr.Code; status == http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}

	// Check the response body is what we expect.
	var body models.LoginResponse
	_ = json.NewDecoder(rr.Body).Decode(&body)
	if token := body.Token; len(token) != 0 {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), token)
	}
}

func TestErrLogin(t *testing.T) {
	rr := makeTest(t,
		"POST",
		"/login",
		"",
		"{\"name\": \"new\",\"password\": \"123\"}",
		Login,
	)

	var body models.LoginResponse
	_ = json.NewDecoder(rr.Body).Decode(&body)

	// Новый пользователь. Должна быть ошибка
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
		t.Errorf("respponse %v", body)
	}

	if len(body.Token) != 0 {
		t.Errorf("handler returned unexpected body: got <%v> want <%v>",
			rr.Body.String(), "Empty")
	} else {
		lastToken = body.Token
	}
}

func TestLogin(t *testing.T) {
	rr := makeTest(t,
		"POST",
		"/login",
		"",
		"{\"name\": \"root\",\"password\": \"123\"}",
		Login,
	)

	var body models.LoginResponse
	_ = json.NewDecoder(rr.Body).Decode(&body)

	// Новый пользователь. Должна быть ошибка
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
		t.Errorf("respponse %v", body)
	}

	if len(body.Token) == 0 {
		t.Errorf("handler returned unexpected body: got <%v> want <%v>",
			rr.Body.String(), "Empty")
	} else {
		lastToken = body.Token
	}
}

func TestLogout(t *testing.T) {
	rr := makeTest(t,
		"GET",
		"/logout",
		lastToken,
		"",
		Logout,
	)

	// Новый пользователь. Должна быть ошибка
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestErrLogout(t *testing.T) {
	rr := makeTest(t,
		"GET",
		"/logout",
		"lastToken",
		"",
		Logout,
	)

	// Новый пользователь. Должна быть ошибка
	if status := rr.Code; status == http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}

func makeTest(t *testing.T, method string, path string, token string, body string, route routeFunc) *httptest.ResponseRecorder {
	req, err := http.NewRequest(method, path, strings.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	if token != "" {
		req.Header.Add("token", token)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(route)

	handler.ServeHTTP(rr, req)
	return rr
}
