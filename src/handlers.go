package src

import (
	"errors"
	"settingsstore/gen/models"
	"settingsstore/gen/restapi/operations"

	"gorm.io/gorm"
)

func (h Handler) Auth(s string) (*models.Principal, error) {
	if s == "test" {
		prin := models.Principal(s)
		return &prin, nil
	}

	return nil, errors.New("Ah ah ah! You didn't say the magic word!")
}

func (h Handler) AddSetting(
	params operations.AddSettingParams,
	principal *models.Principal,
) operations.AddSettingResponder {

	asd := ToEntity(params.Setting)
	err := h.db().Create(asd).Error
	if err != nil {
		return operations.NewAddSettingInternalServerError()
	}

	return operations.NewAddSettingCreated()
}

func (h Handler) GetAllSettings(
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

func (h Handler) GetSetting(
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

func (h Handler) UpdateSetting(
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

func (h Handler) RemoveSetting(
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
