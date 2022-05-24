package repository

import (
	"context"
	"log"

	"../entity"
	"cloud.google.com/go/firestore"
)

type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

// To implement the interface we need to contruct the struct
type repo struct{}

//NewPostRepository
func NewPostRepository() PostRepository {
	return &repo{}
}

const (
	projectId      string = "vib-Ux"
	collectionName string = "posts"
)

func (*repo) Save(post *entity.Post) (*entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		return nil, err
	}
	defer client.Close()
	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"Id":    post.Id,
		"Title": post.Title,
		"Text":  post.Text,
	})
	if err != nil {
		log.Fatalf("Failed Adding a new post: %v", err)
		return nil, err
	}

	return post, nil
}
