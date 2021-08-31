package usecase

import (
	"fmt"

	"github.com/gustavares/hub-api/internal/domain/entity"
	"github.com/gustavares/hub-api/internal/repository"
)

type NoPostsFoundError error

type GetUserPosts struct {
	postRepository repository.PostRepository
}

func NewGetUserPosts(postRepository repository.PostRepository) *GetUserPosts {
	return &GetUserPosts{
		postRepository,
	}
}

func (g *GetUserPosts) Do(userId string) (*[]entity.Post, error) {
	posts, err := g.postRepository.GetByUser(userId)
	if err != nil {
		return nil, err
	}
	if len(*posts) == 0 {
		return nil, NoPostsFoundError(fmt.Errorf("No posts were found for user ID %s.", userId))
	}
	return posts, nil
}
