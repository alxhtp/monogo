package databasehelper

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var (
	ColCreatedAt = "created_at"
	ColUpdatedAt = "updated_at"
	ColDeletedAt = "deleted_at"
	ColID        = "id"
)

func PrepareCreation(tx *gorm.DB) {
	now := time.Now()

	tx.Statement.SetColumn(ColID, uuid.New())
	tx.Statement.SetColumn(ColCreatedAt, now)
	tx.Statement.SetColumn(ColUpdatedAt, now)
}

func PrepareUpdate(tx *gorm.DB) {
	tx.Statement.SetColumn(ColUpdatedAt, time.Now())
}

func PrepareDeletion(tx *gorm.DB) {
	curTime := time.Now()

	tx.Statement.AddClause(clause.Update{})

	tx.Statement.AddClause(clause.Set{
		{Column: clause.Column{Name: ColUpdatedAt}, Value: curTime},
		{Column: clause.Column{Name: ColDeletedAt}, Value: curTime},
	})

	tx.Statement.SetColumn(ColUpdatedAt, curTime)
	tx.Statement.SetColumn(ColDeletedAt, curTime)

	tx.Statement.AddClause(clause.Where{Exprs: []clause.Expression{
		clause.Eq{Column: ColDeletedAt, Value: nil},
	}})

	tx.Statement.Build(
		clause.Update{}.Name(),
		clause.Set{}.Name(),
		clause.Where{}.Name(),
	)
}
