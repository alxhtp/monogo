package dtobase

import "time"

// BasePaginationFilter is a transport-safe pagination filter to avoid import cycles
type BasePaginationFilter struct {
	MinCreated  *time.Time
	MaxCreated  *time.Time
	MinUpdated  *time.Time
	MaxUpdated  *time.Time
	WithDeleted *bool
	ShowCount   *bool
	Offset      *int
	Limit       *int
	OrderBy     *string
}

// BasePaginationResult is a transport-safe pagination output
type BasePaginationResult struct {
	Offset  int
	Limit   int
	OrderBy string
	Count   int
}
