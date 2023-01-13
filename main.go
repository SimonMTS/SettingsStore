package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"settingsstore/api"

	chiMiddleware "github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

//go:generate mkdir -p api
//go:generate oapi-codegen --config=types.conf.yml ./spec.yml
//go:generate oapi-codegen --config=server.conf.yml ./spec.yml

type ApiImplementation struct{}

// TODO enums

func main() {
	fmt.Println("starting...")

	// swagger, err := api.GetSwagger()
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
	// 	os.Exit(1)
	// }

	handler := api.NewStrictHandler(ApiImplementation{}, nil)
	router := chi.NewRouter()

	router.Use(chiMiddleware.Logger)
	// router.Use(middleware.OapiRequestValidator(swagger))
	api.HandlerFromMux(handler, router)

	server := &http.Server{
		Handler: router,
		Addr:    "0.0.0.0:8080",
	}

	fmt.Println("serving...")
	log.Fatal(server.ListenAndServe())
}

func (ai ApiImplementation) GetAllSettings(
	ctx context.Context,
	request api.GetAllSettingsRequestObject,
) (api.GetAllSettingsResponseObject, error) {

	return api.GetAllSettings200JSONResponse{

		api.Setting{
			Name: "test1",
			Rank: 1,
		},
		api.Setting{
			Name: "test2",
			Rank: 2,
		},
	}, nil
}

func (ai ApiImplementation) GetSetting(
	ctx context.Context,
	request api.GetSettingRequestObject,
) (api.GetSettingResponseObject, error) {
	return api.GetSetting200JSONResponse{
		Name: "test name",
		Rank: request.Id,
	}, nil
}

func (ai ApiImplementation) AddSettings(
	ctx context.Context,
	request api.AddSettingsRequestObject,
) (api.AddSettingsResponseObject, error) {
	fmt.Printf("request.Body: %v\n", request.Body)
	return api.AddSettings201Response{}, nil
}

func (ai ApiImplementation) UpdateSettings(
	ctx context.Context,
	request api.UpdateSettingsRequestObject,
) (api.UpdateSettingsResponseObject, error) {
	fmt.Printf("request.Body: %v\n", request.Body)
	return api.UpdateSettings201Response{}, nil
}
