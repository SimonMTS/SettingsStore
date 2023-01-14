package src

import (
	"errors"
	"settingsstore/gen/models"
	"settingsstore/gen/restapi/operations"

	httperr "github.com/go-openapi/errors"

	"gorm.io/gorm"
)

func (h handler) Auth(s string) (*models.Principal, error) {
	if s == "test" {
		prin := models.Principal(s)
		return &prin, nil
	}

	return nil, httperr.New(401, "incorrect api key auth")
}

func (h handler) AddSetting(
	params operations.AddSettingParams,
	principal *models.Principal,
) operations.AddSettingResponder {
	err := h.db().Create(ToEntity(params.Setting)).Error
	if err != nil {
		return operations.NewAddSettingInternalServerError()
	}

	return operations.NewAddSettingCreated()
}

func (h handler) GetAllSettings(
	params operations.GetAllSettingsParams,
	principal *models.Principal,
) operations.GetAllSettingsResponder {

	settingEntities := []Setting{}
	err := h.db().Find(&settingEntities).Error
	if err != nil {
		return operations.NewGetAllSettingsInternalServerError()
	}

	return operations.NewGetAllSettingsOK().
		WithPayload(ToDtos(&settingEntities))
}

func (h handler) GetSetting(
	params operations.GetSettingParams,
	principal *models.Principal,
) operations.GetSettingResponder {

	settingEntity := Setting{ID: params.ID}
	err := h.db().First(&settingEntity).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return operations.NewGetSettingNotFound()
	} else if err != nil {
		return operations.NewGetSettingInternalServerError()
	}

	return operations.NewGetSettingOK().
		WithPayload(ToDto(&settingEntity))
}

func (h handler) UpdateSetting(
	params operations.UpdateSettingParams,
	principal *models.Principal,
) operations.UpdateSettingResponder {

	err := h.db().Save(ToEntity(params.Setting)).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return operations.NewUpdateSettingNotFound()
	} else if err != nil {
		return operations.NewUpdateSettingInternalServerError()
	}

	return operations.NewUpdateSettingNoContent()
}

func (h handler) RemoveSetting(
	params operations.RemoveSettingParams,
	principal *models.Principal,
) operations.RemoveSettingResponder {

	err := h.db().Delete(&Setting{}, params.ID).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return operations.NewRemoveSettingNotFound()
	} else if err != nil {
		return operations.NewRemoveSettingInternalServerError()
	}

	return operations.NewRemoveSettingNoContent()
}
