package src

import (
	"context"
	"fmt"
	"settingsstore/gen/restapi"
	"settingsstore/gen/restapi/operations"
	"settingsstore/gen/restapi/operations/rest"
	"settingsstore/gen/restapi/operations/stream"

	"github.com/MadAppGang/httplog"
	"github.com/go-openapi/loads"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"gorm.io/gorm"
)

type Handler struct {
	Database *gorm.DB
}

func (h Handler) db() *gorm.DB {
	return h.Database.Debug()
}

func ApiSetup(db *gorm.DB) (*restapi.Server, error) {
	fmt.Println("tracing setup...")

	p := tracingSetup()
	otel.SetTracerProvider(p)
	// defer shutdown

	fmt.Println("server setup...")

	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		return nil, err
	}

	api := operations.NewSettingsStoreAPI(swaggerSpec)
	api.UseSwaggerUI()

	handler := Handler{Database: db}
	api.RestAddSettingHandler = rest.AddSettingHandlerFunc(handler.AddSetting)
	api.RestGetAllSettingsHandler = rest.GetAllSettingsHandlerFunc(handler.GetAllSettings)
	api.RestGetSettingHandler = rest.GetSettingHandlerFunc(handler.GetSetting)
	api.StreamSettingUpdatesHandler = stream.SettingUpdatesHandlerFunc(handler.StreamSettings)
	api.RestUpdateSettingHandler = rest.UpdateSettingHandlerFunc(handler.UpdateSetting)
	api.RestRemoveSettingHandler = rest.RemoveSettingHandlerFunc(handler.RemoveSetting)
	api.KeyAuth = handler.Auth

	server := restapi.NewServer(api)
	server.SetHandler((api.Serve(httplog.Logger)))
	server.Port = 8080

	return server, nil
}

func tracingSetup() *trace.TracerProvider {
	ctx := context.Background()

	url := "http://localhost:14268/api/traces"

	r, err := resource.New(ctx, resource.WithAttributes(
		semconv.ServiceNameKey.String("settings-store")))
	if err != nil {
		panic(err)
	}

	e, err := jaeger.New(jaeger.WithCollectorEndpoint(
		jaeger.WithEndpoint(url)))
	if err != nil {
		panic(err)
	}

	b := trace.NewBatchSpanProcessor(e)

	return trace.NewTracerProvider(
		trace.WithSampler(trace.AlwaysSample()),
		trace.WithResource(r),
		trace.WithSpanProcessor(b))
}
