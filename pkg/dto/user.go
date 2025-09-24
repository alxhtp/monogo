package dto

import (
	dtobase "github.com/alxhtp/monogo/pkg/dto/base"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type ReqCreateUser struct {
	Name     string       `json:"name" validate:"required"`
	Email    string       `json:"email" validate:"required,email"`
	Metadata UserMetadata `json:"metadata"`
}

func (r *ReqCreateUser) Validate(validate *validator.Validate) error {
	return validate.Struct(r)
}

type ReqUpdateUser struct {
	Name     *string       `json:"name"`
	Email    *string       `json:"email" validate:"omitempty,email"`
	Status   *int          `json:"status"`
	Metadata *UserMetadata `json:"metadata"`
}

type UserMetadata struct {
	Sex     string `json:"sex" validate:"required,oneof=male female"`
	Address string `json:"address" validate:"required,max=255"`
	Phone   string `json:"phone" validate:"required,e164"`
}

type ReqGetUser struct {
	IDs     *string `query:"ids"` // comma separated string of uuids
	Name    *string `query:"name"`
	Email   *string `query:"email"`
	Status  *int    `query:"status"`
	Sex     *string `query:"sex"`
	Address *string `query:"address"`
	Phone   *string `query:"phone"`
	dtobase.BaseReqQueryPagination
}

type ResUser struct {
	ID       uuid.UUID    `json:"id"`
	Name     string       `json:"name"`
	Email    string       `json:"email"`
	Status   int          `json:"status"`
	Metadata UserMetadata `json:"metadata"`
}

type ResUserSingle struct {
	dtobase.BaseRes
	Data *ResUser `json:"data"`
}

type ResUserList struct {
	dtobase.BaseResPagination
	Data []ResUser `json:"data"`
}
