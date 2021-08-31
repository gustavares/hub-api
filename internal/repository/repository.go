package repository

import "github.com/gustavares/hub-api/internal/domain/entity"

type PostRepository interface {
	// GetById(id string) (entity.Post, error)
	GetByUser(userId string) (*[]entity.Post, error)
}
