package healths

import (
	"time"

	"github.com/hellofresh/health-go/v5"
	healthHttp "github.com/hellofresh/health-go/v5/checks/http"
	healthMongo "github.com/hellofresh/health-go/v5/checks/mongo"
	"github.com/yagoluiz/user-api/internal/config"
)

func NewHealthChecks(cfg *config.Config) *health.Health {
	healths, _ := health.New(health.WithComponent(health.Component{
		Name:    "user-api",
		Version: "v1.0",
	}))

	healths.Register(health.Config{
		Name:      "mongo",
		Timeout:   time.Second * 5,
		SkipOnErr: true,
		Check: healthMongo.New(healthMongo.Config{
			DSN: cfg.MongoConnection,
		}),
	})
	healths.Register(health.Config{
		Name:      "api",
		Timeout:   time.Second * 5,
		SkipOnErr: true,
		Check: healthHttp.New(healthHttp.Config{
			URL: cfg.ApiHost,
		}),
	})

	return healths
}
