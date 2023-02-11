package src

import (
	"encoding/json"
	"errors"
	"net/http"
	"settingsstore/gen/dto"
	"settingsstore/gen/restapi/operations/rest"
	"settingsstore/gen/restapi/operations/stream"
	"time"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (h Handler) Auth(s string) (*dto.Principal, error) {
	if s == "test" {
		prin := dto.Principal(s)
		return &prin, nil
	}

	return nil, errors.New("Ah ah ah! You didn't say the magic word!")
}

func (h Handler) AddSetting(
	params rest.AddSettingParams,
	principal *dto.Principal,
) rest.AddSettingResponder {

	asd := ToEntity(params.Setting)
	err := h.db().Create(asd).Error
	if err != nil {
		return rest.NewAddSettingInternalServerError()
	}

	return rest.NewAddSettingCreated()
}

func (h Handler) GetAllSettings(
	params rest.GetAllSettingsParams,
	principal *dto.Principal,
) rest.GetAllSettingsResponder {

	settingEntities := []Setting{}
	err := h.db().Find(&settingEntities).Error
	if err != nil {
		return rest.NewGetAllSettingsInternalServerError()
	}

	return rest.NewGetAllSettingsOK().
		WithPayload(ToDtos(&settingEntities))
}

func (h Handler) GetSetting(
	params rest.GetSettingParams,
	principal *dto.Principal,
) rest.GetSettingResponder {

	settingEntity := Setting{ID: uuid.MustParse(params.ID.String())}
	err := h.db().First(&settingEntity).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return rest.NewGetSettingNotFound()
	} else if err != nil {
		return rest.NewGetSettingInternalServerError()
	}

	return rest.NewGetSettingOK().
		WithPayload(ToDto(&settingEntity))
}

func (h Handler) StreamSettings(
	params stream.SettingUpdatesParams,
	principal *dto.Principal,
) stream.SettingUpdatesResponder {

	return genericStreamer{h, params.ID}
}

type genericStreamer struct {
	h  Handler
	id strfmt.UUID
}

func (gs genericStreamer) SettingUpdatesResponder() {}
func (gs genericStreamer) WriteResponse(rw http.ResponseWriter, p runtime.Producer) {
	rw.Header().Add("Cache-Control", "no-cache")
	rw.Header().Add("Connection", "keep-alive")
	rw.WriteHeader(200)

	e := json.NewEncoder(rw)
	f, _ := rw.(http.Flusher)

	for i := 0; i < 5; i++ {

		p := dto.Principal("")
		resp := gs.h.GetSetting(rest.GetSettingParams{ID: gs.id}, &p).(*rest.GetSettingOK)

		e.Encode(resp.Payload)
		f.Flush()
		time.Sleep(time.Second)
	}
}

func (h Handler) UpdateSetting(
	params rest.UpdateSettingParams,
	principal *dto.Principal,
) rest.UpdateSettingResponder {

	err := h.db().Save(ToEntity(params.Setting)).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return rest.NewUpdateSettingNotFound()
	} else if err != nil {
		return rest.NewUpdateSettingInternalServerError()
	}

	return rest.NewUpdateSettingNoContent()
}

func (h Handler) RemoveSetting(
	params rest.RemoveSettingParams,
	principal *dto.Principal,
) rest.RemoveSettingResponder {

	err := h.db().Delete(&Setting{}, params.ID).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return rest.NewRemoveSettingNotFound()
	} else if err != nil {
		return rest.NewRemoveSettingInternalServerError()
	}

	return rest.NewRemoveSettingNoContent()
}
