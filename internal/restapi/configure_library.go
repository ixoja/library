// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"database/sql"
	"github.com/ixoja/library/internal/controller"
	"github.com/ixoja/library/internal/handler"
	"github.com/ixoja/library/internal/storage"
	"log"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/ixoja/library/internal/restapi/operations"
)

//go:generate swagger generate server --target ..\..\internal --name Library --spec ..\api\spec.yaml

func configureFlags(api *operations.LibraryAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.LibraryAPI) http.Handler {
	db, err := sql.Open("sqlite3", "./shorten.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Println("failed to start db connection:", err.Error())
		}
	}()

	st := storage.New(db)
	if err := st.InitDB(); err != nil {
		log.Fatalf("failed to init db: %v", err)
	}

	c := controller.New(st)
	h := handler.New(c)

	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.CreateBookHandler == nil {
		api.CreateBookHandler = operations.CreateBookHandlerFunc(func(params operations.CreateBookParams) middleware.Responder {
			return h.CreateBookHandler(params)
		})
	}
	if api.DeleteBookHandler == nil {
		api.DeleteBookHandler = operations.DeleteBookHandlerFunc(func(params operations.DeleteBookParams) middleware.Responder {
			return h.DeleteBookHandler(params)
		})
	}
	if api.GetAllBooksHandler == nil {
		api.GetAllBooksHandler = operations.GetAllBooksHandlerFunc(func(params operations.GetAllBooksParams) middleware.Responder {
			return h.GetAllBooksHandler(params)
		})
	}
	if api.GetBookHandler == nil {
		api.GetBookHandler = operations.GetBookHandlerFunc(func(params operations.GetBookParams) middleware.Responder {
			return h.GetBookHandler(params)
		})
	}
	if api.UpdateBookHandler == nil {
		api.UpdateBookHandler = operations.UpdateBookHandlerFunc(func(params operations.UpdateBookParams) middleware.Responder {
			return h.UpdateBookHandler(params)
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
