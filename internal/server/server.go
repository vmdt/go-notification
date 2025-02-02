package server

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/vmdt/notification-system/config"
	echoserver "github.com/vmdt/notification-system/pkg/http/echo/server"
	"github.com/vmdt/notification-system/pkg/logger"
	"go.uber.org/fx"
)

func RunServer(lc fx.Lifecycle, log logger.ILogger, e *echo.Echo, ctx context.Context, cfg *config.Config) error {
	lc.Append(fx.Hook{
		OnStart: func (_ context.Context) error {
			go func () {
				if err := echoserver.RunHttpServer(ctx, e, log, cfg.Echo); err != nil {
					log.Errorf("error starting http server: %v", err)
				}
			}()

			e.GET("/", func(c echo.Context) error {
				return c.String(200, "Server notification-system is running...")
			})

			return nil
		},
		OnStop: func (_ context.Context) error {
			log.Infof("all servers shutdown gracefully...")
			return nil
		},
	})

	return nil
}