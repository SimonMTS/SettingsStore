package src

import (
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Setting struct {
	ID    uuid.UUID
	Type  string
	Value string
	End   time.Time
}

func DatabaseSetup() (*gorm.DB, error) {
	fmt.Println("database setup...")

	dsn := "host=localhost " +
		"user=postgres " +
		"password=postgres " +
		"dbname=postgres " +
		"port=5432"

	if override := os.Getenv("SETTING_STORE_DB"); override != "" {
		dsn = override
	}

	var (
		db  *gorm.DB
		err error
	)
	for {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}

		time.Sleep(time.Second)
	}

	err = Migrate(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&Setting{})
}
