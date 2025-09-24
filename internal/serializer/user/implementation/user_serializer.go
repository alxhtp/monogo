package userserializerimplementation

import (
	"net/http"

	"github.com/alxhtp/monogo/internal/entity"
	entitybase "github.com/alxhtp/monogo/internal/entity/base"
	userserializer "github.com/alxhtp/monogo/internal/serializer/user"
	"github.com/alxhtp/monogo/pkg/constant"

	"github.com/alxhtp/monogo/pkg/dto"
	dtobase "github.com/alxhtp/monogo/pkg/dto/base"
	databasehelper "github.com/alxhtp/monogo/pkg/helper/database"
	parserhelper "github.com/alxhtp/monogo/pkg/helper/parser"
	queryhelper "github.com/alxhtp/monogo/pkg/helper/query"
)

type userSerializer struct{}

func NewUserSerializer() userserializer.UserSerializer {
	return &userSerializer{}
}

func (s *userSerializer) FilterDTOToEntity(filter dto.ReqGetUser) (entity.UserFilter, error) {
	var (
		output entity.UserFilter
		err    error
	)

	if filter.IDs != nil {
		output.IDs, err = parserhelper.SliceUUIDsStr(*filter.IDs)
		if err != nil {
			return output, err
		}
	}

	if filter.Name != nil {
		output.Name = filter.Name
	}

	if filter.Email != nil {
		output.Email = filter.Email
	}

	if filter.Status != nil {
		status := constant.UserStatus(*filter.Status)
		output.Status = &status
	}

	if filter.Sex != nil {
		output.Sex = filter.Sex
	}

	if filter.Address != nil {
		output.Address = filter.Address
	}

	if filter.Phone != nil {
		output.Phone = filter.Phone
	}

	output.PaginationFilter = queryhelper.SerializeFilterPaginationDtoToEntity(filter.BaseReqQueryPagination)

	return output, err
}

func (s *userSerializer) UpdateDTOToMap(update dto.ReqUpdateUser) (map[string]any, error) {
	var (
		output = make(map[string]any)
		err    error
	)

	if update.Name != nil {
		output["name"] = update.Name
	}

	if update.Email != nil {
		output["email"] = update.Email
	}

	if update.Status != nil {
		output["status"] = update.Status
	}

	if update.Metadata != nil {
		output["metadata"] = databasehelper.GormJsonType[entity.UserMetadata]{
			Item: entity.UserMetadata{
				Sex:     update.Metadata.Sex,
				Address: update.Metadata.Address,
				Phone:   update.Metadata.Phone,
			},
		}
	}

	return output, err
}

func (s *userSerializer) CreateDTOToEntity(create dto.ReqCreateUser) (entity.User, error) {
	var (
		output entity.User
		err    error
	)

	output = entity.User{
		Name:   create.Name,
		Email:  create.Email,
		Status: constant.UserStatusActive,
		Metadata: databasehelper.GormJsonType[entity.UserMetadata]{
			Item: entity.UserMetadata{
				Sex:     create.Metadata.Sex,
				Address: create.Metadata.Address,
				Phone:   create.Metadata.Phone,
			},
		},
	}

	return output, err
}

func (s *userSerializer) EntityToResponse(entity entity.User) dto.ResUser {

	userMetadata := dto.UserMetadata{
		Sex:     entity.Metadata.Item.Sex,
		Address: entity.Metadata.Item.Address,
		Phone:   entity.Metadata.Item.Phone,
	}

	return dto.ResUser{
		ID:       entity.ID,
		Name:     entity.Name,
		Email:    entity.Email,
		Status:   int(entity.Status),
		Metadata: userMetadata,
	}
}

func (s *userSerializer) EntityToResponseSingle(entity *entity.User, code int, message string, stacktrace *string) dto.ResUserSingle {
	var data *dto.ResUser
	if entity != nil {
		res := s.EntityToResponse(*entity)
		data = &res
	}

	isSuccess := code >= http.StatusOK && code < http.StatusMultipleChoices
	return dto.ResUserSingle{
		BaseRes: dtobase.BaseRes{
			Success: isSuccess,
			Code:    code,
			Message: message,
		},
		Data: data,
	}
}

func (s *userSerializer) EntityToResponseList(entities []entity.User, pagination entitybase.BasePaginationResult, code int, message string, stacktrace *string) dto.ResUserList {
	responses := make([]dto.ResUser, len(entities))
	for i, entity := range entities {
		responses[i] = s.EntityToResponse(entity)
	}
	isSuccess := code >= http.StatusOK && code < http.StatusMultipleChoices

	return dto.ResUserList{
		BaseResPagination: dtobase.BaseResPagination{
			BaseRes: dtobase.BaseRes{Code: code, Message: message, Stacktrace: stacktrace, Success: isSuccess},
			Page: dtobase.BasePagination{
				Offset:  pagination.Offset,
				Limit:   pagination.Limit,
				Count:   pagination.Count,
				OrderBy: pagination.OrderBy,
			},
		},
		Data: responses,
	}
}
