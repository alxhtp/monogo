package entitybase

import (
	"time"

	databasehelper "github.com/alxhtp/monogo/pkg/helper/database"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID uuid.UUID `gorm:"column:id;primaryKey;type:uuid;default:uuid_generate_v4()"`
	BaseTime
}

func (b *Base) BeforeCreate(tx *gorm.DB) error {
	databasehelper.PrepareCreation(tx)

	return nil
}

func (b *Base) BeforeUpdate(tx *gorm.DB) error {
	databasehelper.PrepareUpdate(tx)

	return nil
}

func (b *Base) BeforeDelete(tx *gorm.DB) error {
	databasehelper.PrepareDeletion(tx)

	return nil
}

type BaseTime struct {
	CreatedAt time.Time      `gorm:"column:created_at;type:timestamptz;default:now()"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:timestamptz;default:now()"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:timestamptz"`
}
