package userusecase

import (
	"context"

	"github.com/alxhtp/monogo/pkg/dto"
	dtobase "github.com/alxhtp/monogo/pkg/dto/base"
	"github.com/google/uuid"
)

type UserUsecase interface {
	CreateUser(ctx context.Context, req *dto.ReqCreateUser) dto.ResUserSingle
	GetUserByID(ctx context.Context, id uuid.UUID) dto.ResUserSingle
	GetUsersByFilter(ctx context.Context, filter *dto.ReqGetUser) dto.ResUserList
	UpdateUser(ctx context.Context, id uuid.UUID, req *dto.ReqUpdateUser) dto.ResUserSingle
	DeleteUser(ctx context.Context, id uuid.UUID) dtobase.BaseRes
}
