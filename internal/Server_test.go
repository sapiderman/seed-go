package internal

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/sapiderman/seed-go/internal/handlers"
	"github.com/sapiderman/seed-go/internal/router"
	"github.com/sirupsen/logrus"
)

func TestAll(t *testing.T) {

	r := router.NewRouter()
	r.Router = mux.NewRouter()

	logrus.SetLevel(logrus.TraceLevel)

	r.Router = mux.NewRouter()
	h := handlers.NewHealth()
	appRouter.Router.HandleFunc("/health", h.Handler)
	HealthCheckTesting(t)
	NoRouteTesting(t)

	v1 := appRouter.Router.PathPrefix("/v1").Subrouter()
	v1.HandleFunc("/hello", handlers.Hello).Methods("GET")
	v1.HandleFunc("/time", handlers.GetTime).Methods("GET")

	HelloHandlerTesting(t)
	TimeHandlerTesting(t)
}

func HealthCheckTesting(t *testing.T) {
	t.Log("Testing HealthCheck")
	recorder := httptest.NewRecorder()
	healthRequest := httptest.NewRequest("GET", "/health", nil)
	appRouter.Router.ServeHTTP(recorder, healthRequest)
	if recorder.Code != http.StatusOK {
		t.Errorf("expecting healthcheck status 200 but %d", recorder.Code)
		t.FailNow()
	}
}

func NoRouteTesting(t *testing.T) {

	t.Log("Testing nonexistent route")
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/nonexistent", nil)
	appRouter.Router.ServeHTTP(recorder, request)
	if recorder.Code == http.StatusOK {
		t.Errorf("expecting http status 404 but got %d", recorder.Code)
		t.FailNow()
	}
}

func HelloHandlerTesting(t *testing.T) {

	t.Log("Testing /v1/hello")
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/v1/hello", nil)
	appRouter.Router.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusOK {
		t.Errorf("expecting status 200 but got %d", recorder.Code)
		t.FailNow()
	}

}

func TimeHandlerTesting(t *testing.T) {

	t.Log("Testing /v1/time")
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/v1/time", nil)
	appRouter.Router.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusOK {
		t.Errorf("expecting http status 200 but got %d", recorder.Code)
		t.FailNow()
	}

}
