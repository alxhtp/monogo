package userserializer

import (
	"github.com/alxhtp/monogo/internal/entity"
	entitybase "github.com/alxhtp/monogo/internal/entity/base"
	"github.com/alxhtp/monogo/pkg/dto"
)

type UserSerializer interface {
	FilterDTOToEntity(filter dto.ReqGetUser) (entity.UserFilter, error)
	UpdateDTOToMap(update dto.ReqUpdateUser) (map[string]any, error)
	CreateDTOToEntity(create dto.ReqCreateUser) (entity.User, error)

	EntityToResponse(entity entity.User) dto.ResUser
	EntityToResponseSingle(entity *entity.User, code int, message string, stacktrace *string) dto.ResUserSingle
	EntityToResponseList(entities []entity.User, pagination entitybase.BasePaginationResult, code int, message string, stacktrace *string) dto.ResUserList
}
