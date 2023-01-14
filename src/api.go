package main

import (
	"settingsstore/gen/models"
	"settingsstore/gen/restapi/operations"

	"gorm.io/gorm"
)

var settings []*models.Setting

type handler struct {
	db *gorm.DB
}

func (h handler) AddSettings(params operations.AddSettingParams) operations.AddSettingResponder {
	h.db.Create(ToEntity(params.Setting))
	return operations.NewAddSettingCreated()
}

func (h handler) GetAllSettings(params operations.GetAllSettingsParams) operations.GetAllSettingsResponder {
	settingEntities := []SettingEntity{}
	h.db.Find(&settingEntities)
	return operations.NewGetAllSettingsOK().WithPayload(ToDtos(&settingEntities))
}

func (h handler) GetSetting(params operations.GetSettingParams) operations.GetSettingResponder {
	settingEntity := SettingEntity{ID: params.ID}
	h.db.First(&settingEntity)
	return operations.NewGetSettingOK().WithPayload(ToDto(&settingEntity))
}

func (h handler) UpdateSettings(params operations.UpdateSettingParams) operations.UpdateSettingResponder {
	h.db.Save(ToEntity(params.Setting))
	return operations.NewUpdateSettingNoContent()
}

func (h handler) RemoveSettings(params operations.RemoveSettingParams) operations.RemoveSettingResponder {
	h.db.Delete(&SettingEntity{}, params.ID)
	return operations.NewRemoveSettingNoContent()
}
