package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	const port string = ":8000"
	router.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Println(resp, "Up and Running..")
	})
	router.HandleFunc("/posts", getPosts).Methods("GET")
	router.HandleFunc("/posts", addPost).Methods("POST")
	log.Println("Server Listening on port: ", port)
	log.Fatalln(http.ListenAndServe(port, router))
}
