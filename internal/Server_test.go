package internal

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/sapiderman/seed-go/internal/handlers"
	"github.com/sirupsen/logrus"
)

func pretifyJSON(sjson string) string {
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(sjson), &m)
	if err != nil {
		return "mot a json"
	}
	byt, err := json.MarshalIndent(m, "", "   ")
	if err != nil {
		return err.Error()
	}
	return string(byt)
}

func TestAll(t *testing.T) {
	var srv Server

	logrus.SetLevel(logrus.TraceLevel)

	srv.Router = mux.NewRouter()
	h := handlers.NewHealth()
	srv.Router.HandleFunc("/health", h.Handler)
	srv.HealthCheckTesting(t)
	srv.NoRouteTesting(t)

	v1 := srv.Router.PathPrefix("/v1").Subrouter()
	v1.HandleFunc("/hello", handlers.HandlerHello).Methods("GET")
	v1.HandleFunc("/time", handlers.HandlerGetTime).Methods("GET")

	srv.HelloHandlerTesting(t)
	srv.TimeHandlerTesting(t)
}

func (s *Server) HealthCheckTesting(t *testing.T) {
	t.Log("Testing HealthCheck")
	recorder := httptest.NewRecorder()
	healthRequest := httptest.NewRequest("GET", "/health", nil)
	s.Router.ServeHTTP(recorder, healthRequest)
	if recorder.Code != http.StatusOK {
		t.Errorf("expecting healthcheck status 200 but %d", recorder.Code)
		t.FailNow()
	}
}

func (s *Server) NoRouteTesting(t *testing.T) {

	t.Log("Testing nonexistent route")
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/nonexistent", nil)
	s.Router.ServeHTTP(recorder, request)
	if recorder.Code == http.StatusOK {
		t.Errorf("expecting http status 404 but got %d", recorder.Code)
		t.FailNow()
	}
}

func (s *Server) HelloHandlerTesting(t *testing.T) {

	t.Log("Testing /v1/hello")
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/v1/hello", nil)
	s.Router.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusOK {
		t.Errorf("expecting status 200 but got %d", recorder.Code)
		t.FailNow()
	}

}

func (s *Server) TimeHandlerTesting(t *testing.T) {

	t.Log("Testing /v1/time")
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/v1/time", nil)
	s.Router.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusOK {
		t.Errorf("expecting http status 200 but got %d", recorder.Code)
		t.FailNow()
	}

}
