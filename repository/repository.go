package repository

import (
	"context"

	"github.com/akamiko/entity-sample2/model"
)

type UserRepository interface {
	GetByID(ctx context.Context, id int) (*model.User, error)
}
