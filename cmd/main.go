package main

import (
	"net/http"

	"github.com/evry-bergen/toolbox-reference-golang/pkg/api"

	"github.com/gorilla/mux"
)

func setupRouter() *mux.Router {
	router := mux.NewRouter()

	router.Get("/").HandlerFunc(api.GetHandler)

	return router
}

func main() {
	r := setupRouter()
	srv := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: r,
	}

	srv.ListenAndServe()
}
