package rest

import (
	"wash-payment/internal/app"
	"wash-payment/internal/config"
	"wash-payment/internal/pkg/openapi/restapi"
	"wash-payment/internal/pkg/openapi/restapi/operations"
	"wash-payment/internal/pkg/openapi/restapi/operations/standard"

	"github.com/go-openapi/loads"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type service struct {
	l *zap.SugaredLogger
}

func NewServer(cfg config.Config, l *zap.SugaredLogger) (*restapi.Server, error) {
	service := &service{
		l: l,
	}

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load embedded swagger spec")
	}

	api := operations.NewWashPaymentAPI(swaggerSpec)

	api.Logger = service.l.Infof

	api.StandardHealthCheckHandler = standard.HealthCheckHandlerFunc(healthCheck)

	server := restapi.NewServer(api)
	server.Host = cfg.Host
	server.Port = cfg.Port

	return server, nil
}

func healthCheck(params standard.HealthCheckParams, profile *app.Auth) standard.HealthCheckResponder {
	return standard.NewHealthCheckOK().WithPayload(&standard.HealthCheckOKBody{Ok: true})
}
