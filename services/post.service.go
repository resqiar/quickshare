package services

import (
	"quickshare/entities"
	"quickshare/inputs"
	"quickshare/repositories"
)

type PostService interface {
	CreatePost(input *inputs.CreatePostInput, userId string) (string, error)
	FindPostByID(id string) (*entities.Post, error)
}

type PostServiceImpl struct {
	Repository repositories.PostRepo
}

func (s *PostServiceImpl) CreatePost(input *inputs.CreatePostInput, userId string) (string, error) {
	result, err := s.Repository.CreatePost(input, userId)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (s *PostServiceImpl) FindPostByID(id string) (*entities.Post, error) {
	result, err := s.Repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return result, nil
}
