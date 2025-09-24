package userrepository

import (
	"context"

	"github.com/alxhtp/monogo/internal/entity"
	entitybase "github.com/alxhtp/monogo/internal/entity/base"
	"github.com/google/uuid"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) (output *entity.User, err error)
	GetByID(ctx context.Context, id uuid.UUID) (output *entity.User, err error)
	GetByFilter(ctx context.Context, filter *entity.UserFilter) (output []entity.User, paginationResult entitybase.BasePaginationResult, err error)
	Update(ctx context.Context, id uuid.UUID, updateMap map[string]any) (output *entity.User, err error)
	Delete(ctx context.Context, id uuid.UUID) (err error)
}
