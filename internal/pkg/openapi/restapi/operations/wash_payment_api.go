// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"wash-payment/internal/app/entity"
	"wash-payment/internal/pkg/openapi/restapi/operations/organizations"
	"wash-payment/internal/pkg/openapi/restapi/operations/standard"
)

// NewWashPaymentAPI creates a new WashPayment instance
func NewWashPaymentAPI(spec *loads.Document) *WashPaymentAPI {
	return &WashPaymentAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),

		JSONProducer: runtime.JSONProducer(),

		OrganizationsDepositHandler: organizations.DepositHandlerFunc(func(params organizations.DepositParams, principal *entity.Auth) organizations.DepositResponder {
			return organizations.DepositNotImplemented()
		}),
		OrganizationsGetHandler: organizations.GetHandlerFunc(func(params organizations.GetParams, principal *entity.Auth) organizations.GetResponder {
			return organizations.GetNotImplemented()
		}),
		StandardHealthCheckHandler: standard.HealthCheckHandlerFunc(func(params standard.HealthCheckParams) standard.HealthCheckResponder {
			return standard.HealthCheckNotImplemented()
		}),
		OrganizationsListHandler: organizations.ListHandlerFunc(func(params organizations.ListParams, principal *entity.Auth) organizations.ListResponder {
			return organizations.ListNotImplemented()
		}),
		OrganizationsSetServicePricesHandler: organizations.SetServicePricesHandlerFunc(func(params organizations.SetServicePricesParams, principal *entity.Auth) organizations.SetServicePricesResponder {
			return organizations.SetServicePricesNotImplemented()
		}),
		OrganizationsTransactionsHandler: organizations.TransactionsHandlerFunc(func(params organizations.TransactionsParams, principal *entity.Auth) organizations.TransactionsResponder {
			return organizations.TransactionsNotImplemented()
		}),

		// Applies when the "Authorization" header is set
		AuthKeyAuth: func(token string) (*entity.Auth, error) {
			return nil, errors.NotImplemented("api key auth (authKey) Authorization from header param [Authorization] has not yet been implemented")
		},
		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*WashPaymentAPI Payment system service for self-service car washes */
type WashPaymentAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator

	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator

	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// AuthKeyAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key Authorization provided in the header
	AuthKeyAuth func(string) (*entity.Auth, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// OrganizationsDepositHandler sets the operation handler for the deposit operation
	OrganizationsDepositHandler organizations.DepositHandler
	// OrganizationsGetHandler sets the operation handler for the get operation
	OrganizationsGetHandler organizations.GetHandler
	// StandardHealthCheckHandler sets the operation handler for the health check operation
	StandardHealthCheckHandler standard.HealthCheckHandler
	// OrganizationsListHandler sets the operation handler for the list operation
	OrganizationsListHandler organizations.ListHandler
	// OrganizationsSetServicePricesHandler sets the operation handler for the set service prices operation
	OrganizationsSetServicePricesHandler organizations.SetServicePricesHandler
	// OrganizationsTransactionsHandler sets the operation handler for the transactions operation
	OrganizationsTransactionsHandler organizations.TransactionsHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *WashPaymentAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *WashPaymentAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *WashPaymentAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *WashPaymentAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *WashPaymentAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *WashPaymentAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *WashPaymentAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *WashPaymentAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *WashPaymentAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the WashPaymentAPI
func (o *WashPaymentAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.AuthKeyAuth == nil {
		unregistered = append(unregistered, "AuthorizationAuth")
	}

	if o.OrganizationsDepositHandler == nil {
		unregistered = append(unregistered, "organizations.DepositHandler")
	}
	if o.OrganizationsGetHandler == nil {
		unregistered = append(unregistered, "organizations.GetHandler")
	}
	if o.StandardHealthCheckHandler == nil {
		unregistered = append(unregistered, "standard.HealthCheckHandler")
	}
	if o.OrganizationsListHandler == nil {
		unregistered = append(unregistered, "organizations.ListHandler")
	}
	if o.OrganizationsSetServicePricesHandler == nil {
		unregistered = append(unregistered, "organizations.SetServicePricesHandler")
	}
	if o.OrganizationsTransactionsHandler == nil {
		unregistered = append(unregistered, "organizations.TransactionsHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *WashPaymentAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *WashPaymentAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	result := make(map[string]runtime.Authenticator)
	for name := range schemes {
		switch name {
		case "authKey":
			scheme := schemes[name]
			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, func(token string) (interface{}, error) {
				return o.AuthKeyAuth(token)
			})

		}
	}
	return result
}

// Authorizer returns the registered authorizer
func (o *WashPaymentAPI) Authorizer() runtime.Authorizer {
	return o.APIAuthorizer
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *WashPaymentAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *WashPaymentAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *WashPaymentAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the wash payment API
func (o *WashPaymentAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *WashPaymentAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/organizations/{id}/deposit"] = organizations.NewDeposit(o.context, o.OrganizationsDepositHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/organizations/{id}"] = organizations.NewGet(o.context, o.OrganizationsGetHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/healthCheck"] = standard.NewHealthCheck(o.context, o.StandardHealthCheckHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/organizations"] = organizations.NewList(o.context, o.OrganizationsListHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/organizations/{id}/service-prices"] = organizations.NewSetServicePrices(o.context, o.OrganizationsSetServicePricesHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/organizations/{id}/transactions"] = organizations.NewTransactions(o.context, o.OrganizationsTransactionsHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *WashPaymentAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *WashPaymentAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *WashPaymentAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *WashPaymentAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *WashPaymentAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[um][path] = builder(h)
	}
}
