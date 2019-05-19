package main

import (
	"log"
	"net/http"
	"github.com/googollee/go-socket.io"
)

func main() {
	server, err := socketio.NewServer(nil)

	if err != nil {
		log.Fatal(err)		
	}

	// sockets 
	server.On("connection", func (so socketio.Socket) {
		log.Println("A new user connected")

		// other events
	})

	// http
	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./public")))
	log.Println("Serving at localhost:8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))

}