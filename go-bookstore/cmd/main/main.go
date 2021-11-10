/*
	In the main file we will first create the server
	call the routes

*/
package main

import (
	"log"
	"net/http"

	"github.com/Vib-UX/go-bookstore/pkg/routes" // In go we generally deal with absolute paths
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r)) // ListenAndServe creates server and we pass the port number
}
