package dtobase

import (
	"time"
)

// BaseRes envelope struct for non paginated response
type BaseRes struct {
	Success    bool    `json:"success" validate:"required"`
	Code       int     `json:"code" validate:"required"`
	Message    string  `json:"message" validate:"required"`
	Stacktrace *string `json:"stacktrace,omitempty"`
}

func (b *BaseRes) Error() string {
	return b.Message
}

// BasePagination base page struct
type BasePagination struct {
	Offset  int    `json:"offset" validate:"required"`
	Limit   int    `json:"limit" validate:"required"`
	Count   int    `json:"count" validate:"required"`
	OrderBy string `json:"order_by" validate:"required"`
}

// BaseResPagination envelope struct for paginated response
type BaseResPagination struct {
	BaseRes
	Page BasePagination `json:"page"`
}

// BaseResTime envelope struct for entity timestamps
type BaseResTime struct {
	CreatedAt time.Time  `json:"created_at" validate:"required"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
