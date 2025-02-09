package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/vmdt/notification-system/config"
	"github.com/vmdt/notification-system/internal/server"
	"github.com/vmdt/notification-system/pkg/http"
	echoserver "github.com/vmdt/notification-system/pkg/http/echo/server"
	"github.com/vmdt/notification-system/pkg/logger"
	gorm_postgres "github.com/vmdt/notification-system/pkg/postgres"
	"github.com/vmdt/notification-system/pkg/rabbitmq"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Options(
			fx.Provide(
				config.InitConfig,
				logger.InitLogger,
				rabbitmq.NewRabbitMQConn,
				echoserver.NewEchoServer,
				http.NewContext,
				validator.New,
				gorm_postgres.NewGorm,
			),
			fx.Invoke(server.RunServer),
		),
	).Run()
}