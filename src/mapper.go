package src

import (
	"settingsstore/gen/models"

	"github.com/google/uuid"
)

func ToEntity(source *models.Setting) (target *Setting) {
	target = &Setting{}
	target.ID = UuidToEntity(source.ID)
	target.Type = *source.Type
	target.Value = *source.Value
	target.End = source.End.UTC()
	return
}

func ToDto(source *Setting) (target *models.Setting) {
	target = &models.Setting{}
	target.ID = UuidToDto(source.ID)
	target.Type = &source.Type
	target.Value = &source.Value
	target.End = &models.DateTime{Time: source.End}
	return
}

func ToDtos(sources *[]Setting) (targets []*models.Setting) {
	for _, s := range *sources {
		targets = append(targets, ToDto(&s))
	}
	return
}

func UuidToEntity(source *models.UUID) uuid.UUID {
	return uuid.MustParse(source.String())
}

func UuidToDto(source uuid.UUID) *models.UUID {
	return &models.UUID{
		UUID: source,
	}
}
