package internal

import (
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
func (r *Router) InitRoutes() {

	fmt.Println("initializing routes")

	r.MuxRouter.HandleFunc("/health", handlers.HandlerHealth).Methods("GET")
	r.MuxRouter.HandleFunc("/hello", handlers.HandlerHello).Methods("GET")

}
