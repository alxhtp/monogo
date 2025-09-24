package entitybase

import (
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"
)

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

type BasePaginationResult struct {
	Offset  int
	Limit   int
	OrderBy string
	Count   int
}

// GenerateBaseOrderMap generate base order map
// to disable default order keys, override the key with empty string value
func GenerateBaseOrderMap() map[string]bool {
	return map[string]bool{
		"created_at": true,
		"updated_at": true,
		"deleted_at": true,
	}
}

const (
	defaultLimit         = 100
	maxLimit             = 1000
	defaultGormQuerySort = ".created_at desc"
)

// PaginateWithLimit ...
func PaginateWithLimit(f *BasePaginationFilter, minLimit, maxLimit int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if f == nil {
			return db
		}
		if f.Offset != nil {
			db = db.Offset(*f.Offset)
		}
		if f.Limit != nil {
			db = db.Limit(*f.Limit)
		}
		return db
	}
}

// PaginateEntityQuery do gorm pagination filter and orders based on pagination
// baseTableName should not from client request, must from constant. No string escape guard.
// paginationResult and filter should be provided before function call
func PaginateEntityQuery(
	db *gorm.DB,
	baseTableName string,
	orderMap map[string]bool,
	filter *BasePaginationFilter,
	paginationResult *BasePaginationResult,
) *gorm.DB {
	if db == nil || filter == nil || paginationResult == nil {
		return db
	}

	if filter.WithDeleted != nil && *filter.WithDeleted {
		db = db.Unscoped()
	}

	if filter.MinCreated != nil {
		db = db.Where(baseTableName+".created_at >= ?", *filter.MinCreated)
	}

	if filter.MaxCreated != nil {
		db = db.Where(baseTableName+".created_at <= ?", *filter.MaxCreated)
	}

	if filter.MinUpdated != nil {
		db = db.Where(baseTableName+".updated_at >= ?", *filter.MinUpdated)
	}

	if filter.MaxUpdated != nil {
		db = db.Where(baseTableName+".updated_at <= ?", *filter.MaxUpdated)
	}

	var count int64
	db.Count(&count)
	paginationResult.Count = int(count)

	if filter.Offset != nil && *filter.Offset > 0 {
		db = db.Offset(*filter.Offset)
		paginationResult.Offset = *filter.Offset
	}

	limit := defaultLimit
	if filter.Limit != nil && *filter.Limit > 0 && *filter.Limit <= maxLimit {
		limit = *filter.Limit
	}
	db = db.Limit(limit)
	paginationResult.Limit = limit

	if filter.OrderBy != nil {
		paginationResult.OrderBy = *filter.OrderBy
		return OrderEntityQuery(db, *filter.OrderBy, orderMap)
	}

	return db.Order(baseTableName + defaultGormQuerySort)
}

// OrderEntityQuery implement order query param into order query db statements
func OrderEntityQuery(db *gorm.DB, orderByQueryParam string, orderMap map[string]bool) *gorm.DB {
	orderStatements := OrderQueryTranslator(orderByQueryParam, orderMap)

	for i := range orderStatements {
		var orderStatement = orderStatements[i]
		db = db.Order(orderStatement)
	}

	return db
}

// OrderQueryTranslator translate every word into database order statements
func OrderQueryTranslator(orderByQueryParam string, orderMap map[string]bool) []string {
	out := make([]string, 0)

	if orderMap == nil || len(orderByQueryParam) == 0 {
		return out
	}

	words := strings.Split(orderByQueryParam, ",")

	for i := range words {
		var (
			word      = strings.TrimSpace(words[i])
			direction = "asc"
		)

		if len(word) == 0 {
			continue
		}

		if word[0:1] == "+" {
			word = word[1:]
		}

		if word[0:1] == "-" {
			direction = "desc"
			word = word[1:]
		}

		if enableOrder, ok := orderMap[word]; ok && enableOrder {
			out = append(out, fmt.Sprintf(`"%s" %s`, word, direction))
		}
	}

	return out
}
