package databasehelper

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/alxhtp/monogo/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dbOnceByKey = map[string]*sync.Once{}
	dbByKey     = map[string]*gorm.DB{}
)

const (
	defaultConnName   = "default"
	defaultSearchPath = "public"
)

func NewGormDB(ctx context.Context, cfg *config.DatabaseConfig) (*gorm.DB, error) {
	// Compose key for singleton map. In future, these can be passed/derived.
	connName := defaultConnName
	searchPath := defaultSearchPath
	key := connName + "-" + searchPath

	// Fast path: reuse if healthy
	if cached := getDBForKey(key); cached != nil {
		if sqlDB, err := cached.DB(); err == nil {
			if err := sqlDB.PingContext(ctx); err == nil {
				return cached, nil
			}
		}
		// fallthrough to reconnect if ping fails
	}

	// Ensure once exists for this key
	once := ensureOnceForKey(key)

	var openErr error
	once.Do(func() {
		var db *gorm.DB
		db, openErr = openGormWithConfig(cfg)
		if openErr != nil {
			return
		}
		// Validate connection
		if sqlDB, err := db.DB(); err != nil {
			openErr = err
			return
		} else if err := sqlDB.PingContext(ctx); err != nil {
			openErr = err
			return
		}
		setDBForKey(key, db)
	})
	if openErr != nil {
		return nil, openErr
	}
	db := getDBForKey(key)
	return db, nil
}

func openGormWithConfig(cfg *config.DatabaseConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBSSLMode,
	)
	db, err := gorm.Open(postgres.New(postgres.Config{DSN: dsn, PreferSimpleProtocol: true}), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	// Apply pool settings
	if cfg.MaxIdleConns > 0 {
		sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	}
	if cfg.MaxOpenConns > 0 {
		sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	}
	if cfg.ConnMaxLifetime != "" {
		if d, err := time.ParseDuration(cfg.ConnMaxLifetime); err == nil {
			sqlDB.SetConnMaxLifetime(d)
		}
	}
	return db, nil
}

// Storage and helpers for singleton db connections
var dbMu sync.Mutex

func ensureOnceForKey(key string) *sync.Once {
	dbMu.Lock()
	defer dbMu.Unlock()
	if o, ok := dbOnceByKey[key]; ok {
		return o
	}
	o := &sync.Once{}
	dbOnceByKey[key] = o
	return o
}

func getDBForKey(key string) *gorm.DB {
	dbMu.Lock()
	defer dbMu.Unlock()
	return dbByKey[key]
}

func setDBForKey(key string, db *gorm.DB) {
	dbMu.Lock()
	defer dbMu.Unlock()
	dbByKey[key] = db
}
