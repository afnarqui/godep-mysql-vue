package main

import (
	"encoding/json"
	"fmt"
	"net/http"	
	"github.com/go-chi/chi"
)

func main() {
	port := ":8082"
	r := chi.NewRouter()

	fmt.Println("Serving on " + port)
	http.ListenAndServe(port, r)
}