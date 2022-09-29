/*
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package router

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/NuVeS/URLShort/cmd/apis/auth"
	"github.com/NuVeS/URLShort/cmd/apis/urls"
	"github.com/NuVeS/URLShort/cmd/logger"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = logger.Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	router.NotFoundHandler = http.HandlerFunc(urls.Route)

	router.Use()
	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/index",
		Index,
	},

	Route{
		"Login",
		strings.ToUpper("Post"),
		"/login",
		auth.Login,
	},

	Route{
		"Logout",
		strings.ToUpper("Get"),
		"/logout",
		auth.Logout,
	},

	Route{
		"Register",
		strings.ToUpper("Post"),
		"/register",
		auth.Register,
	},

	Route{
		"Delete",
		strings.ToUpper("Post"),
		"/delete",
		urls.Delete,
	},

	Route{
		"List",
		strings.ToUpper("Get"),
		"/list",
		urls.List,
	},

	Route{
		"Shorten",
		strings.ToUpper("Post"),
		"/shorten",
		urls.Shorten,
	},
}
