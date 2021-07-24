package router

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/sapiderman/seed-go/api"
	"github.com/sapiderman/seed-go/internal/handlers"
	"github.com/sapiderman/seed-go/internal/middlewares"

	log "github.com/sirupsen/logrus"
)

// Router is a wrapper for all the router connections
type Router struct {
	Router   *mux.Router        // point to mux routers
	Handlers *handlers.Handlers // point to handlers
}

var rLog = log.WithField("module", "router")

// NewRouter get new Instance
func NewRouter() *Router {
	return &Router{}
}

// InitRoutes creates our routes
func InitRoutes(router *Router) {
	rLog.WithField("fn", "InitRoutes()").Info("Initializing routes...")

	r := router.Router
	rh := router.Handlers

	// register middlewares
	// r.Use(apmgorilla.Middleware()) // apmgorilla.Instrument(r.MuxRouter) // elastic apm
	r.Use(middlewares.ContextStart, middlewares.MyLogger, middlewares.ContextEnd) // ye-olde middlewares

	// health check endpoint. Not in a version path as it will seems to be a permanent endpoint (famous last words)
	h := handlers.NewHealth()
	r.HandleFunc("/health", h.Handler)

	// v1 APIs
	v1 := r.PathPrefix("/v1").Subrouter()
	v1.HandleFunc("/hello", handlers.Hello).Methods("GET")
	v1.HandleFunc("/time", handlers.GetTime).Methods("GET")

	v1.HandleFunc("/users", rh.ListUsers).Methods("GET")
	v1.HandleFunc("/devices", rh.ListDevices).Methods("GET")
	v1.HandleFunc("/user", rh.AddUser).Methods("POST")
	v1.HandleFunc("/device", rh.AddDevice).Methods("POST")

	// static file handler
	r.PathPrefix("/web/").Handler(http.StripPrefix("/web/", http.FileServer(http.Dir("./web"))))

	// handle swagger api static files as /docs.
	// r.HandleFunc("/docs", api.ServeStatic).Methods("GET")
	for path := range api.StaticResources {
		r.HandleFunc(path, api.ServeStatic).Methods("GET")
	}

	// v2 as pgx APIs
	// v2 := r.PathPrefix("/v2").Subrouter()
	// v2.HandleFunc("/devices", handlers.PgxListDevices).Methods("GET")

	// r.NotFoundHandler = http.HandlerFunc(handlers.HandlerNotFound)
	// r.HandleFunc("/", handlers.HandlerNotFound)
	// http.Handle("/", http.FileServer(http.Dir("./static/404.html")))

	// display routes
	walk(*r)
}

// walk runs the mux.Router.Walk method to print all the registerd routes
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
		log.Error(err)
	}
}
