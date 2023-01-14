package main

import (
	"fmt"
	"log"
	"settingsstore/gen/restapi"
	"settingsstore/gen/restapi/operations"

	"github.com/go-openapi/loads"
)

//go:generate rm -rf gen
//go:generate mkdir -p gen
//go:generate swagger generate server -t gen -f ./spec.yml --exclude-main --strict-responders
//go:generate go mod tidy

func main() {
	fmt.Println("starting...")

	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewSettingsStoreAPI(swaggerSpec)
	api.AddSettingHandler = operations.AddSettingHandlerFunc(AddSettings)
	api.GetAllSettingsHandler = operations.GetAllSettingsHandlerFunc(GetAllSettings)
	api.GetSettingHandler = operations.GetSettingHandlerFunc(GetSetting)
	api.UpdateSettingHandler = operations.UpdateSettingHandlerFunc(UpdateSettings)
	api.RemoveSettingHandler = operations.RemoveSettingHandlerFunc(RemoveSettings)

	server := restapi.NewServer(api)
	defer server.Shutdown()
	server.Port = 8080
	server.ConfigureAPI()

	fmt.Println("Interactive docs on : http://localhost:8080/docs")
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
