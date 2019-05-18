package main

import (
	"net/http"
	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome afn port 8082"))
	})
	http.ListenAndServe(":8082", r)
}