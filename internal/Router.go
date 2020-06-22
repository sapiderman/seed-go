package internal

import (
	"context"
	"fmt"

	"github.com/gorilla/mux"
	"github.com/sapiderman/test-seed/internal/handlers"
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

	r.MuxRouter.HandleFunc("/health", handlers.HandlerHealth).Methods("GET")

	// v1 APIs
	v1 := r.MuxRouter.PathPrefix("/v1").Subrouter()
	v1.HandleFunc("/hello", handlers.HandlerHello).Methods("GET")

}
