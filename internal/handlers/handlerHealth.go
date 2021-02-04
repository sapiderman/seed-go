package handlers

import (
	"github.com/nelkinda/health-go"
	"github.com/nelkinda/health-go/checks/sysinfo"
	"github.com/nelkinda/health-go/checks/uptime"
)

// NewHealth returns a new instance of health service
func NewHealth() *health.Service {

	return health.New(
		health.Health{Version: "1", ReleaseID: "1.0.0-SNAPSHOT"},
		uptime.System(),
		uptime.Process(),
		sysinfo.Health(),
	)
}
