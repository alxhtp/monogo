package databasehelper

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	jsonconvert "github.com/alxhtp/monogo/pkg/jsonconvert"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

type GormJsonType[T any] struct {
	Item T
}

func (ia *GormJsonType[T]) Scan(value interface{}) error {
	var (
		err   error
		bytes []byte
	)

	if value == nil {
		return nil
	}

	if bytes, err = ParseToBytes(value); err != nil {
		return err
	}

	var result T
	if err = json.Unmarshal(bytes, &result); err != nil {
		return err
	}

	*ia = GormJsonType[T]{
		Item: result,
	}

	return nil
}

// String representation
func (ia GormJsonType[T]) String() string {
	return jsonconvert.Serialize(ia.Item)
}

// GormDataType gorm common data type
func (ia GormJsonType[T]) GormDataType() string {
	return jsonconvert.JsonGormDataType()
}

// GormDBDataType gorm db data type
func (ia GormJsonType[T]) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	return jsonconvert.JsonGormDBDataType(db, field)
}

// GormValue gorm value
func (ia GormJsonType[T]) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	return jsonconvert.JsonGormValue(ctx, ia.Item, db)
}

func ParseToBytes(value interface{}) ([]byte, error) {
	var bytes []byte
	switch v := value.(type) {
	case []byte:
		if len(v) > 0 {
			bytes = make([]byte, len(v))
			copy(bytes, v)
		}
	case string:
		bytes = []byte(v)
	default:
		return nil, errors.New(fmt.Sprint("failed convert to []byte:", value))
	}

	return bytes, nil
}
