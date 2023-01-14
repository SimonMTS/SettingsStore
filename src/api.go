package src

import (
	"fmt"
	"settingsstore/gen/restapi"
	"settingsstore/gen/restapi/operations"

	"github.com/go-openapi/loads"
	"gorm.io/gorm"
)

func ApiSetup(db *gorm.DB) (*restapi.Server, error) {
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

type handler struct {
	db *gorm.DB
}

func (h handler) AddSettings(
	params operations.AddSettingParams,
) operations.AddSettingResponder {
	h.db.Create(ToEntity(params.Setting))
	return operations.NewAddSettingCreated()
}

func (h handler) GetAllSettings(
	params operations.GetAllSettingsParams,
) operations.GetAllSettingsResponder {
	settingEntities := []SettingEntity{}
	h.db.Find(&settingEntities)
	return operations.NewGetAllSettingsOK().
		WithPayload(ToDtos(&settingEntities))
}

func (h handler) GetSetting(
	params operations.GetSettingParams,
) operations.GetSettingResponder {
	settingEntity := SettingEntity{ID: params.ID}
	h.db.First(&settingEntity)
	return operations.NewGetSettingOK().
		WithPayload(ToDto(&settingEntity))
}

func (h handler) UpdateSettings(
	params operations.UpdateSettingParams,
) operations.UpdateSettingResponder {
	h.db.Save(ToEntity(params.Setting))
	return operations.NewUpdateSettingNoContent()
}

func (h handler) RemoveSettings(
	params operations.RemoveSettingParams,
) operations.RemoveSettingResponder {
	h.db.Delete(&SettingEntity{}, params.ID)
	return operations.NewRemoveSettingNoContent()
}
