package dtobase

import "time"

// BaseReqQueryPagination base pagination filter query dto
type BaseReqQueryPagination struct {
	CreatedAtGTE   *time.Time `query:"created-at-gte"`
	CreatedAtLTE   *time.Time `query:"created-at-lte"`
	UpdatedAtGTE   *time.Time `query:"updated-at-gte"`
	UpdatedAtLTE   *time.Time `query:"updated-at-lte"`
	IncludeDeleted *bool      `query:"include-deleted"`
	ShowCount      *bool      `query:"show-count"`
	Offset         *int       `query:"offset"`
	Limit          *int       `query:"limit"`
	OrderBy        *string    `query:"order-by"`
}
