package src

import (
	"fmt"
	"settingsstore/gen/restapi"
	"settingsstore/gen/restapi/operations"
	"settingsstore/gen/restapi/operations/rest"
	"settingsstore/gen/restapi/operations/stream"

	"github.com/MadAppGang/httplog"
	"github.com/go-openapi/loads"
	"gorm.io/gorm"
)

type Handler struct {
	Database *gorm.DB
}

func (h Handler) db() *gorm.DB {
	return h.Database.Debug()
}

func ApiSetup(db *gorm.DB) (*restapi.Server, error) {
	fmt.Println("server setup...")

	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		return nil, err
	}

	api := operations.NewSettingsStoreAPI(swaggerSpec)
	api.UseSwaggerUI()

	handler := Handler{Database: db}
	api.RestAddSettingHandler = rest.AddSettingHandlerFunc(handler.AddSetting)
	api.RestGetAllSettingsHandler = rest.GetAllSettingsHandlerFunc(handler.GetAllSettings)
	api.RestGetSettingHandler = rest.GetSettingHandlerFunc(handler.GetSetting)
	api.StreamSettingUpdatesHandler = stream.SettingUpdatesHandlerFunc(handler.StreamSettings)
	api.RestUpdateSettingHandler = rest.UpdateSettingHandlerFunc(handler.UpdateSetting)
	api.RestRemoveSettingHandler = rest.RemoveSettingHandlerFunc(handler.RemoveSetting)
	api.KeyAuth = handler.Auth

	server := restapi.NewServer(api)
	server.SetHandler((api.Serve(httplog.Logger)))
	server.Port = 8080

	return server, nil
}
