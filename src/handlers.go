package src

import (
	"encoding/json"
	"errors"
	"net/http"
	"settingsstore/gen/models"
	"settingsstore/gen/restapi/operations/rest"
	"settingsstore/gen/restapi/operations/stream"
	"time"

	"github.com/go-openapi/runtime"
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
	params rest.AddSettingParams,
	principal *models.Principal,
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
	principal *models.Principal,
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
	principal *models.Principal,
) rest.GetSettingResponder {

	settingEntity := Setting{ID: params.ID}
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
	principal *models.Principal,
) stream.SettingUpdatesResponder {

	return genericStreamer{h}
}

type genericStreamer struct{ h Handler }

func (gs genericStreamer) SettingUpdatesResponder() {}
func (gs genericStreamer) WriteResponse(rw http.ResponseWriter, p runtime.Producer) {
	rw.Header().Add("Cache-Control", "no-cache")
	rw.Header().Add("Connection", "keep-alive")
	rw.WriteHeader(200)

	e := json.NewEncoder(rw)
	f, _ := rw.(http.Flusher)

	for i := 0; i < 5; i++ {

		p := models.Principal("")
		resp := gs.h.GetSetting(rest.GetSettingParams{ID: 1}, &p).(*rest.GetSettingOK)

		e.Encode(resp.Payload)
		f.Flush()
		time.Sleep(time.Second)
	}
}

func (h Handler) UpdateSetting(
	params rest.UpdateSettingParams,
	principal *models.Principal,
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
	principal *models.Principal,
) rest.RemoveSettingResponder {

	err := h.db().Delete(&Setting{}, params.ID).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return rest.NewRemoveSettingNotFound()
	} else if err != nil {
		return rest.NewRemoveSettingInternalServerError()
	}

	return rest.NewRemoveSettingNoContent()
}
