package usecase

import (
	"github.com/yk2220s/go-grpc-sample/api/domain/entity"

	"errors"
)

// IGetPost is usecase to get a post
type IGetPost interface {
	Execute(id int) (*entity.Post, error)
}

// GetPost implements IGetPost
type GetPost struct{}

// Execute get a post by id
func (uc *GetPost) Execute(id int) (*entity.Post, error) {
	if id == 1 {
		return &entity.Post{
			ID:    1,
			Title: "How to train the dragon",
			Text:  "the amazing film",
		}, nil
	}

	return nil, errors.New("NotFound: Post")
}

// NewGetPost is provider
func NewGetPost() IGetPost {
	return &GetPost{}
}
