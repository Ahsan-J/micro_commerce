package router

import (
	"cart/graphql"
	"cart/responses"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GenerateRoutes generates the routes by resolves route and controllers import
func GenerateRoutes(router *mux.Router) {
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		responses.GetCartSuccess.SendAPI(w, "Hello")
	})

	router.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		result := graphql.ExecuteQuery(r.URL.Query().Get("query"))
		json.NewEncoder(w).Encode(result)
	})
}
