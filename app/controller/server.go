package controller

import (
	"fmt"
	"log"
	"net/http"

	controller "github.com/Pedr0visk/go-simple-rest-api/app/controller/blog"
	"github.com/gorilla/mux"
)

var router *mux.Router

func initHandlers() {
	// we could have /api/post/{id} for GET/PUT/DELETE
	// router.HandleFunc("/api/post/{id}", controller.{appropriateMethod}).Methods("{GET or POST or PUT")
	// instead of what we have.  So the endpoint would be the same to read, update, and delete
	// we'd just have different handlers for those actions
	router.HandleFunc("/api/posts", controller.GetAllPosts).Methods("GET")
	router.HandleFunc("/api/post/{id}", controller.GetPost).Methods("GET")
	router.HandleFunc("/api/post/new", controller.CreatePost).Methods("POST")
	router.HandleFunc("/api/post/update", controller.UpdatePost).Methods("PUT")
	router.HandleFunc("/api/post/delete/{id}", controller.DeletePost).Methods("DELETE")
}

func Start() {
	router = mux.NewRouter()

	initHandlers()
	fmt.Printf("router initialized and listening on 3200\n")
	log.Fatal(http.ListenAndServe(":3200", router))
}
