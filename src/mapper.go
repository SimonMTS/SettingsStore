package main

import (
	"settingsstore/gen/models"
)

func ToEntity(source *models.Setting) (target *SettingEntity) {
	target = &SettingEntity{}
	target.ID = *source.ID
	target.Type = *source.Type
	target.Value = *source.Value
	target.End = source.End.UTC()
	return
}

func ToDto(source *SettingEntity) (target *models.Setting) {
	target = &models.Setting{}
	target.ID = &source.ID
	target.Type = &source.Type
	target.Value = &source.Value
	target.End = &models.DateTime{Time: source.End}
	return
}

func ToDtos(sources *[]SettingEntity) (targets []*models.Setting) {
	for _, s := range *sources {
		targets = append(targets, ToDto(&s))
	}
	return
}
