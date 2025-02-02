package rabbitmq

import (
	"context"
	"fmt"
	"time"

	"github.com/cenkalti/backoff/v4"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type RabbitMQConfig struct {
	Host     string
	Port     int
	User     string
	Password string
}

func NewRabbitMQConn(cfg *RabbitMQConfig, ctx context.Context) (*amqp.Connection, error) {
	connAddr := fmt.Sprintf(
		"amqp://%s:%s@%s:%d",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
	)

	bo := backoff.NewExponentialBackOff()
	bo.MaxElapsedTime = 10 * time.Second // Maximum time to retry
	maxRetries := 5 // Number of retries

	var conn *amqp.Connection
	var err error

	err = backoff.Retry(func () error {
		conn, err = amqp.Dial(connAddr)
		if err != nil {
			log.Errorf("Failed to connect to RabbitMQ: %v", err)
			return err
		}

		return nil
	}, backoff.WithMaxRetries(bo, uint64(maxRetries - 1)))

	log.Info("Connected to RabbitMQ")
	
	go func() {
		select {
		case <-ctx.Done():
			err := conn.Close()
			if err != nil {
				log.Errorf("Failed to close RabbitMQ connection: %v", err)
			}

			log.Info("Closed RabbitMQ connection")
		}
	}()

	return conn, err
}