package src

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Setting struct {
	ID    int64
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

	err = db.AutoMigrate(&Setting{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
