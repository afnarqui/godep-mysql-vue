package main

import (
	"httpd/handler"
	"platform/newsfeed"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	port := ":8082"
	feed := newsfeed.New()

	r := chi.NewRouter()

	r.Get("/newsfeed", handler.NewsfeedGet(feed))
	r.Post("/newsfeed", handler.NewsfeedPost(feed))

	fmt.Println("Serving on " + port)
	http.ListenAndServe(port, r)
}
