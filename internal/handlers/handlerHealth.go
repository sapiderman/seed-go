package handlers

import (
	"context"

	"github.com/nelkinda/health-go"
	"github.com/nelkinda/health-go/checks/sysinfo"
	"github.com/nelkinda/health-go/checks/uptime"
)

// HealthResponse is just the resp to return
// type healthResponse struct {
// 	ServerStatus     string `json:"serverStatus"`
// 	ServerTime       string `json:"serverTime"`
// 	ServerUpDuration uint64 `json:"serverUpDuration"`

// 	ServerVersion string `json:"serverVersion"`
// }

// func HandlerHealth(w http.ResponseWriter, s *http.Request) {} : DISABLED, using health-go lib below

// NewHealth returns a new instance of health ervice
func NewHealth(ctx context.Context) *health.Service {

	return health.New(
		health.Health{Version: "1", ReleaseID: "1.0.0-SNAPSHOT"},
		uptime.System(),
		uptime.Process(),
		sysinfo.Health(),
	)
}
