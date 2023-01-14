package src

import (
	"fmt"
	"settingsstore/gen/restapi"
	"settingsstore/gen/restapi/operations"

	"github.com/MadAppGang/httplog"
	"github.com/go-openapi/loads"
	"gorm.io/gorm"
)

type handler struct {
	database *gorm.DB
}

func (h handler) db() *gorm.DB {
	return h.database.Debug()
}

func ApiSetup(db *gorm.DB) (*restapi.Server, error) {
	fmt.Println("server setup...")

	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		return nil, err
	}

	api := operations.NewSettingsStoreAPI(swaggerSpec)
	api.UseSwaggerUI()

	handler := handler{database: db}
	api.AddSettingHandler = operations.AddSettingHandlerFunc(handler.AddSetting)
	api.GetAllSettingsHandler = operations.GetAllSettingsHandlerFunc(handler.GetAllSettings)
	api.GetSettingHandler = operations.GetSettingHandlerFunc(handler.GetSetting)
	api.UpdateSettingHandler = operations.UpdateSettingHandlerFunc(handler.UpdateSetting)
	api.RemoveSettingHandler = operations.RemoveSettingHandlerFunc(handler.RemoveSetting)
	api.KeyAuth = handler.Auth

	server := restapi.NewServer(api)
	server.SetHandler((api.Serve(httplog.Logger)))
	server.Port = 8080

	return server, nil
}
