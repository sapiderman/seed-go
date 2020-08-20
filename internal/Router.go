package internal

import (
	"context"
	"fmt"
	"strings"

	"github.com/gorilla/mux"
	"github.com/sapiderman/seed-go/internal/handlers"
	"github.com/sapiderman/seed-go/internal/logger"

	"go.elastic.co/apm/module/apmgorilla"
)

// Router stores the Mux instance.
type Router struct {
	MuxRouter *mux.Router
}

// NewRouter instantiates and returns new Router
func NewRouter() *Router {

	return &Router{
		MuxRouter: mux.NewRouter(),
	}

}

// InitRoutes creates our routes
func (r *Router) InitRoutes(ctx context.Context) {

	fmt.Println("initializing routes")
	server := ctx.Value(ContextKey(ServerKey)).(*Server)

	// middlewares
	r.MuxRouter.Use(apmgorilla.Middleware()) //	apmgorilla.Instrument(r.MuxRouter) // elastic apm
	r.MuxRouter.Use(logger.MyLogger)         // ye-olde logger

	// health check endpoint. Not in a version path as it will seems to be a permanent endpoint (famous last words)
	r.MuxRouter.HandleFunc("/health", handlers.HandlerHealth).Methods("GET")
	// handle swagger api static files as /docs.
	r.MuxRouter.PathPrefix("/docs").Handler(server.StaticFilter)

	// v1 APIs
	v1 := r.MuxRouter.PathPrefix("/v1").Subrouter()
	v1.HandleFunc("/hello", handlers.HandlerHello).Methods("GET")

}

// walk runs the mux.Router.Walk method to print all the registerd router.
func walk(r mux.Router) {
	err := r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		pathTemplate, err := route.GetPathTemplate()
		if err == nil {
			fmt.Println("ROUTE:", pathTemplate)
		}
		pathRegexp, err := route.GetPathRegexp()
		if err == nil {
			fmt.Println("Path regexp:", pathRegexp)
		}
		queriesTemplates, err := route.GetQueriesTemplates()
		if err == nil {
			fmt.Println("Queries templates:", strings.Join(queriesTemplates, ","))
		}
		queriesRegexps, err := route.GetQueriesRegexp()
		if err == nil {
			fmt.Println("Queries regexps:", strings.Join(queriesRegexps, ","))
		}
		methods, err := route.GetMethods()
		if err == nil {
			fmt.Println("Methods:", strings.Join(methods, ","))
		}
		fmt.Println()
		return nil
	})

	if err != nil {
		fmt.Println(err)
	}
}
