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
	h.db.Create(&SettingEntity{
		ID:    *params.Setting.ID,
		Type:  *params.Setting.Type,
		Value: *params.Setting.Value,
		End:   params.Setting.End.UTC(),
	})

	return operations.NewAddSettingCreated()
}

func (h handler) GetAllSettings(params operations.GetAllSettingsParams) operations.GetAllSettingsResponder {
	settingEntities := []SettingEntity{}
	h.db.Find(&settingEntities)

	settings := []*models.Setting{}
	for _, s := range settingEntities {
		copy := s
		settings = append(settings, &models.Setting{
			ID:    &copy.ID,
			Type:  &copy.Type,
			Value: &copy.Value,
			End:   &models.DateTime{Time: copy.End},
		})
	}

	return operations.NewGetAllSettingsOK().WithPayload(settings)
}

func (h handler) GetSetting(params operations.GetSettingParams) operations.GetSettingResponder {
	settingEntity := SettingEntity{ID: params.ID}
	h.db.First(&settingEntity)
	setting := &models.Setting{
		ID:    &settingEntity.ID,
		Type:  &settingEntity.Type,
		Value: &settingEntity.Value,
		End:   &models.DateTime{Time: settingEntity.End},
	}

	return operations.NewGetSettingOK().WithPayload(setting)
}

func (h handler) UpdateSettings(params operations.UpdateSettingParams) operations.UpdateSettingResponder {
	settingEntity := SettingEntity{ID: params.ID}
	h.db.First(&settingEntity)
	settingEntity.ID = *params.Setting.ID
	settingEntity.Type = *params.Setting.Type
	settingEntity.Value = *params.Setting.Value
	settingEntity.End = params.Setting.End.UTC()
	h.db.Save(&settingEntity)

	return operations.NewUpdateSettingNoContent()
}

func (h handler) RemoveSettings(params operations.RemoveSettingParams) operations.RemoveSettingResponder {
	settingEntity := SettingEntity{ID: params.ID}
	h.db.Delete(&settingEntity)

	return operations.NewRemoveSettingNoContent()
}
