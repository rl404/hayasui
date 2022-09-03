package repository

import (
	"context"

	"github.com/rl404/hayasui/internal/domain/reaction/entity"
)

// Repository contains functions for reaction domain.
type Repository interface {
	SetCommand(ctx context.Context, msgID string, data entity.Command) error
	GetCommand(ctx context.Context, msgID string) (*entity.Command, error)
}
