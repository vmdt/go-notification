package echoserver

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/vmdt/notification-system/pkg/logger"
)

const (
	MaxHeaderBytes = 1 << 20
	ReadTimeout    = 15 * time.Second
	WriteTimeout   = 15 * time.Second
)

type EchoConfig struct {
	Port          string   `mapstructure:"port" validate:"required"`
	Host          string   `mapstructure:"host"`
	IgnoreLogURLs []string `mapstructure:"ignoreLogURLs"`
	BasePath      string   `mapstructure:"basePath" validate:"required"`
	Timeout       int      `mapstructure:"timeout"`
}

func NewEchoServer() *echo.Echo {
	return echo.New()
}

func RunHttpServer(ctx context.Context, echo *echo.Echo, log logger.ILogger, cfg *EchoConfig) error {
	echo.Server.ReadTimeout = ReadTimeout
	echo.Server.WriteTimeout = WriteTimeout
	echo.Server.MaxHeaderBytes = MaxHeaderBytes

	go func () {
		for {
			select {
			case <-ctx.Done():
				log.Infof("Shutting down Http Port: {%s}", cfg.Port)
				err := echo.Server.Shutdown(ctx)
				if err != nil {
					log.Errorf("(Shutdown) error: {%v}", err)
					return
				}
				log.Info("server exited properly")
				return
			}
		}
	}()

	err := echo.Start(cfg.Port)
	return err
}