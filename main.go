package main

import (
	"encoding/json"
	"net/http"
    "/go/src/github.com/afnarqui/src/system/app/newsfeed.go"

	"github.com/go-chi/chi"
)

func main() {
	port := ":8082"
	feed := newsfeed.New()
	feed.Add(newsfeed.Item{
		Title: "Hello",
		Post:  "World",
	})
	r := chi.NewRouter()
	r.Get("/newsfeed", func(w http.ResponseWriter, r *http.Request) {
		items := feed.GetAll()
		json.NewEncoder(w).Encode(items)
	})
	http.ListenAndServe(":8083", r)
}
