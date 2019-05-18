package main

import (
	"encoding/json"
	"fmt"
	"net/http"	
	"github.com/go-chi/chi"
)

func NewsfeedGet(feed newsfeed.Getter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		items := feed.GetAll()
		json.NewEncoder(w).Encode(items)
	}
}

func NewsfeedPost(feed newsfeed.Adder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := map[string]string{}
		json.NewDecoder(r.Body).Decode(&request)

		feed.Add(newsfeed.Item{
			Title: request["title"],
			Post:  request["post"],
		})

		w.Write([]byte("Good job!"))
	}
}

type Getter interface {
	GetAll() []Item
}

type Adder interface {
	Add(item Item)
}

type Item struct {
	Title string `json:"title"`
	Post  string `json:"post"`
}

type Repo struct {
	Items []Item
}

func New() *Repo {
	return &Repo{
		Items: []Item{},
	}
}

func (r *Repo) Add(item Item) {
	r.Items = append(r.Items, item)
}

func (r *Repo) GetAll() []Item {
	return r.Items
}

func main() {
	port := ":8082"
	feed := newsfeed.New()

	r := chi.NewRouter()

	r.Get("/newsfeed", handler.NewsfeedGet(feed))
	r.Post("/newsfeed", handler.NewsfeedPost(feed))

	fmt.Println("Serving on " + port)
	http.ListenAndServe(port, r)
}