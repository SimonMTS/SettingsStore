package src

import (
	"settingsstore/gen/dto"

	"github.com/google/uuid"
)

func ToEntity(source *dto.Setting) (target *Setting) {
	target = &Setting{}
	target.ID = UuidToEntity(source.ID)
	target.Type = *source.Type
	target.Value = *source.Value
	target.End = source.End.UTC()
	return
}

func ToDto(source *Setting) (target *dto.Setting) {
	target = &dto.Setting{}
	target.ID = UuidToDto(source.ID)
	target.Type = &source.Type
	target.Value = &source.Value
	target.End = &dto.DateTime{Time: source.End}
	return
}

func ToDtos(sources *[]Setting) (targets []*dto.Setting) {
	for _, s := range *sources {
		targets = append(targets, ToDto(&s))
	}
	return
}

func UuidToEntity(source *dto.UUID) uuid.UUID {
	return uuid.MustParse(source.String())
}

func UuidToDto(source uuid.UUID) *dto.UUID {
	return &dto.UUID{
		UUID: source,
	}
}
