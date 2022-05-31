package main

import (
	"encoding/json"
	"math/rand"
	"net/http"

	"./entity"
	"./repository"
)

var repo repository.PostRepository = repository.NewFirestoreRepository()

func getPosts(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	posts, err := repo.FindAll()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error":"Error getting the data from repo"}`))
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(posts)
}

func addPost(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var post entity.Post
	err := json.NewDecoder(request.Body).Decode(&post)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error":"Error unmarshalling data"}`))
		return
	}
	post.Id = rand.Int63()
	repo.Save(&post)
	response.WriteHeader(http.StatusOK)
	result, err := json.Marshal(post)
	response.Write(result)
	json.NewEncoder(response).Encode(post)

}
