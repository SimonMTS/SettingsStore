package src

import (
	"fmt"
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
		"port=5432 "

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
