package entity

import (
	entitybase "github.com/alxhtp/monogo/internal/entity/base"
	"github.com/alxhtp/monogo/pkg/constant"
	databasehelper "github.com/alxhtp/monogo/pkg/helper/database"
	"github.com/google/uuid"
)

type User struct {
	entitybase.Base
	Name     string                                    `gorm:"column:name;type:varchar(255);not null"`
	Email    string                                    `gorm:"column:email;type:varchar(255);not null;unique"`
	Status   constant.UserStatus                       `gorm:"column:status;type:int;not null;default:0"`
	Metadata databasehelper.GormJsonType[UserMetadata] `gorm:"column:metadata;type:jsonb"`
}

type UserMetadata struct {
	Sex     string `json:"sex"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

type UserFilter struct {
	IDs              []uuid.UUID
	Name             *string
	Email            *string
	Status           *constant.UserStatus
	Sex              *string
	Address          *string
	Phone            *string
	PaginationFilter entitybase.BasePaginationFilter
}

func (u *User) TableName() string {
	return "monogo.users"
}

func (u *User) OrderMap() map[string]bool {
	out := entitybase.GenerateBaseOrderMap()

	out["name"] = true
	out["email"] = true

	return out
}
