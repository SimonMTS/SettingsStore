package main

import (
	"fmt"
	"log"
	"settingsstore/gen/restapi"
	"settingsstore/gen/restapi/operations"

	"github.com/go-openapi/loads"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//go:generate rm -rf ../gen
//go:generate mkdir -p ../gen
//go:generate swagger generate server -t ../gen -f ../spec.yml --exclude-main --strict-responders
//go:generate go mod tidy

func main() {
	db, err := databaseSetup()
	if err != nil {
		log.Fatalln(err)
	}

	server, err := apiSetup(db)
	if err != nil {
		log.Fatalln(err)
	}

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}

func databaseSetup() (*gorm.DB, error) {
	fmt.Println("database setup...")

	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Europe/Amsterdam"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&SettingEntity{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func apiSetup(db *gorm.DB) (*restapi.Server, error) {
	fmt.Println("server setup...")

	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		return nil, err
	}

	handlers := handler{db: db}
	api := operations.NewSettingsStoreAPI(swaggerSpec)
	api.AddSettingHandler = operations.AddSettingHandlerFunc(handlers.AddSettings)
	api.GetAllSettingsHandler = operations.GetAllSettingsHandlerFunc(handlers.GetAllSettings)
	api.GetSettingHandler = operations.GetSettingHandlerFunc(handlers.GetSetting)
	api.UpdateSettingHandler = operations.UpdateSettingHandlerFunc(handlers.UpdateSettings)
	api.RemoveSettingHandler = operations.RemoveSettingHandlerFunc(handlers.RemoveSettings)

	server := restapi.NewServer(api)
	server.Port = 8080
	server.ConfigureAPI()
	return server, nil
}
