package internal

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

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
	logrus.SetLevel(logrus.TraceLevel)

	r := NewRouter()

	r.HealthCheckTesting(t)
}

func (r *Router) HealthCheckTesting(t *testing.T) {
	t.Log("Testing HealthCheck")
	recorder := httptest.NewRecorder()
	healthRequest := httptest.NewRequest("GET", "/health", nil)
	r.MuxRouter.ServeHTTP(recorder, healthRequest)
	if recorder.Code != http.StatusOK {
		t.Errorf("expecting healthcheck status 200 but %d", recorder.Code)
		t.FailNow()
	} else {
		t.Log(pretifyJSON(recorder.Body.String()))
	}
}
