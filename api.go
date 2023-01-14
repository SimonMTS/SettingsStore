package main

import (
	"settingsstore/gen/models"
	"settingsstore/gen/restapi/operations"
)

var settings []*models.Setting

func AddSettings(params operations.AddSettingParams) operations.AddSettingResponder {
	settings = append(settings, params.Setting)
	return operations.NewAddSettingCreated()
}

func GetAllSettings(params operations.GetAllSettingsParams) operations.GetAllSettingsResponder {
	return operations.NewGetAllSettingsOK().WithPayload(settings)
}

func GetSetting(params operations.GetSettingParams) operations.GetSettingResponder {
	for _, setting := range settings {
		if *setting.ID == params.ID {
			return operations.NewGetSettingOK().WithPayload(setting)
		}
	}

	return &operations.GetSettingNotFound{}
}

func UpdateSettings(params operations.UpdateSettingParams) operations.UpdateSettingResponder {
	for index, setting := range settings {
		if *setting.ID == params.ID {
			settings[index] = params.Setting
			return &operations.UpdateSettingNoContent{}
		}
	}

	return &operations.UpdateSettingNotFound{}
}

func RemoveSettings(params operations.RemoveSettingParams) operations.RemoveSettingResponder {
	for index, setting := range settings {
		if *setting.ID == params.ID {
			settings = append(settings[:index], settings[index+1:]...)
			return &operations.RemoveSettingNoContent{}
		}
	}

	return &operations.RemoveSettingNotFound{}
}
