package main

import (
	"encoding/json"
	"net/http"
)

type Post struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

var posts []Post

func init() {
	posts = []Post{{1, "title 1", "text 1"}}
}

func getPosts(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	result, err := json.Marshal(posts)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error":"Error marshalling data"}`))
	}
	response.WriteHeader(http.StatusOK)
	response.Write(result)
}

func addPost(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var post Post
	err := json.NewDecoder(request.Body).Decode(&post)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error":"Error unmarshalling data"}`))
		return
	}
	post.Id = len(posts) + 1
	posts = append(posts, post)
	response.WriteHeader(http.StatusOK)
	result, err := json.Marshal(post)
	response.Write(result)

}
