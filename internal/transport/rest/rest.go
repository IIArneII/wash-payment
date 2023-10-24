package rest

import (
	"net/http"
	"path"
	"strings"
	"wash-payment/internal/app"
	"wash-payment/internal/config"
	"wash-payment/internal/pkg/openapi/restapi"
	"wash-payment/internal/pkg/openapi/restapi/operations"
	"wash-payment/internal/pkg/openapi/restapi/operations/standard"
	"wash-payment/internal/transport/firebase"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/pkg/errors"
	"github.com/powerman/structlog"
	"github.com/rs/cors"
	"github.com/sebest/xff"
	"go.uber.org/zap"
)

type service struct {
	l        *zap.SugaredLogger
	auth     firebase.FirebaseService
	services *app.Services
}

func NewServer(l *zap.SugaredLogger, cfg config.Config, services *app.Services, auth firebase.FirebaseService) (*restapi.Server, error) {
	svc := &service{
		l:        l,
		auth:     auth,
		services: services,
	}

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to load embedded swagger spec")
	}
	api := operations.NewWashPaymentAPI(swaggerSpec)

	l.Info("SwaggerSpec initialized")

	api.Logger = svc.l.Infof
	api.AuthKeyAuth = svc.auth.Auth

	api.StandardHealthCheckHandler = standard.HealthCheckHandlerFunc(svc.healthCheck)
	svc.initOrganizationsHandlers(api)

	l.Info("Hendlers initialized")

	server := restapi.NewServer(api)
	server.Host = cfg.Host
	server.Port = cfg.Port

	globalMiddlewares := func(handler http.Handler) http.Handler {
		xffmw, _ := xff.Default()
		logger := makeLogger(cfg.BasePath, svc.l)
		accesslog := makeAccessLog(cfg.BasePath, svc.l)
		redocOpts := middleware.RedocOpts{
			BasePath: cfg.BasePath,
			SpecURL:  path.Join(cfg.BasePath, "/swagger.json"),
		}
		return xffmw.Handler(
			logger(
				noCache(
					recovery(
						accesslog(
							middleware.Spec(cfg.BasePath, restapi.FlatSwaggerJSON, middleware.Redoc(redocOpts, handler))), svc.l))))
	}

	l.Info("Middlewares initialized")

	newCORS := cors.New(cors.Options{
		AllowedOrigins:   splitCommaSeparatedStr(cfg.AllowedOrigins),
		AllowedMethods:   []string{"POST", "PUT", "PATCH", "GET", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		Debug:            true,
	})
	newCORS.Log = cors.Logger(structlog.New(structlog.KeyUnit, "CORS"))
	handleCORS := newCORS.Handler

	l.Info("CORS initialized")

	server.SetHandler(handleCORS(api.Serve(globalMiddlewares)))

	return server, nil
}

func (svc *service) healthCheck(params standard.HealthCheckParams, profile *app.Auth) standard.HealthCheckResponder {
	svc.l.Info(profile.User.ID)
	return standard.NewHealthCheckOK().WithPayload(&standard.HealthCheckOKBody{Ok: true})
}

func splitCommaSeparatedStr(str string) (result []string) {
	for _, item := range strings.Split(str, ",") {
		item = strings.TrimSpace(item)
		if item != "" {
			result = append(result, item)
		}
	}
	return
}
