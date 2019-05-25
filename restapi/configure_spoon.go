// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/xbcsmith/spoon/restapi/operations"
	"github.com/xbcsmith/spoon/restapi/operations/spoons"

	handlers "github.com/xbcsmith/spoon/handlers"
	models "github.com/xbcsmith/spoon/models"
)

//go:generate swagger generate server --target ../../spoon --name Spoon --spec ../swagger.yml --principal models.Principal

func configureFlags(api *operations.SpoonAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.SpoonAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "x-spoon-token" header is set
	api.KeyAuth = func(token string) (*models.Principal, error) {
		principal, err := handlers.TokenAuth(token)
		if err != nil {
			api.Logger("Access attempt with incorrect api key auth: %s", token)
			return nil, errors.New(401, "incorrect api key auth")
		}
		return principal, nil
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()
	if api.SpoonsAddSpoonHandler == nil {
		api.SpoonsAddSpoonHandler = spoons.AddSpoonHandlerFunc(func(params spoons.AddSpoonParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation spoons.AddSpoon has not yet been implemented")
		})
	}
	if api.SpoonsDestroySpoonHandler == nil {
		api.SpoonsDestroySpoonHandler = spoons.DestroySpoonHandlerFunc(func(params spoons.DestroySpoonParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation spoons.DestroySpoon has not yet been implemented")
		})
	}
	if api.SpoonsFindSpoonsHandler == nil {
		api.SpoonsFindSpoonsHandler = spoons.FindSpoonsHandlerFunc(func(params spoons.FindSpoonsParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation spoons.FindSpoons has not yet been implemented")
		})
	}
	if api.SpoonsGetSpoonHandler == nil {
		api.SpoonsGetSpoonHandler = spoons.GetSpoonHandlerFunc(func(params spoons.GetSpoonParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation spoons.GetSpoon has not yet been implemented")
		})
	}
	if api.SpoonsUpdateSpoonHandler == nil {
		api.SpoonsUpdateSpoonHandler = spoons.UpdateSpoonHandlerFunc(func(params spoons.UpdateSpoonParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation spoons.UpdateSpoon has not yet been implemented")
		})
	}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
