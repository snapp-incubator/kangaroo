package config

import (
	"github.com/snapp-incubator/kangaroo/internal/infra/logger"
	"github.com/snapp-incubator/kangaroo/internal/infra/telemetry"
	"go.uber.org/fx"
)

// Default return default configuration.
func Default() Config {
	return Config{
		Out: fx.Out{},
		Logger: logger.Config{
			Level: "debug",
		},
		Database: db.Config{
			Name: "koochooloo",
			URL:  "mongodb://127.0.0.1:27017",
		},
		Telemetry: telemetry.Config{
			Namespace:   "1995parham.me",
			ServiceName: "koochooloo",
			Meter: telemetry.Meter{
				Address: ":8080",
				Enabled: true,
			},
			Trace: telemetry.Trace{
				Enabled: false,
				Agent: telemetry.Agent{
					Port: "6831",
					Host: "127.0.0.1",
				},
			},
		},
	}
}
