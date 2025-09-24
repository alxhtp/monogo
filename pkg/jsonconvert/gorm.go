package jsonconvert

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

// JsonGormDataType gorm common data type
func JsonGormDataType() string {
	return "json"
}

// JsonGormDBDataType gorm db data type for JSON
func JsonGormDBDataType(db *gorm.DB, field *schema.Field) string {
	switch db.Name() {
	case "sqlite":
		return "JSON"
	case "mysql":
		return "JSON"
	case "postgres":
		return "JSONB"
	}
	return ""
}

// JsonGormValue gorm JSON value representation
func JsonGormValue(ctx context.Context, obj interface{}, db *gorm.DB) clause.Expr {
	if obj == nil {
		return gorm.Expr("NULL")
	}

	str := Serialize(obj)

	switch db.Name() {
	case "mysql":
		return gorm.Expr("CAST(? AS JSON)", str)
	}

	return gorm.Expr("?", str)
}
