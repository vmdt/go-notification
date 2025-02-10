package rabbitmq

import (
	"context"

	jsoniter "github.com/json-iterator/go"
	"github.com/streadway/amqp"
	"github.com/vmdt/notification-system/pkg/logger"
)

type IPublisher interface {
	PublishMessage(msg interface{}) error
}

type Publisher struct {
	cfg  *RabbitMQConfig
	conn *amqp.Connection
	log  logger.ILogger
	ctx context.Context
}

func NewPublisher(cfg *RabbitMQConfig, conn *amqp.Connection, log logger.ILogger, ctx context.Context) *Publisher {
	return &Publisher{
		cfg:  cfg,
		conn: conn,
		log:  log,
		ctx: ctx,
	}
}

func (p Publisher) PublishMessage(msg interface{}, exchange string, kind string, key string) error {
	data, err := jsoniter.Marshal(msg)
	if err != nil {
		p.log.Error("Error in marshalling message to publish message")
	}

	ch, err := p.conn.Channel()
	if err != nil {
		p.log.Error("Error in getting channel to publish message")
		return err
	}

	defer ch.Close()


	if err != nil {
		p.log.Error("Error in declaring exchange to publish message")
		return err
	}

	publishsingMsg := amqp.Publishing{
		Body: data,
		ContentType: "application/json",
		DeliveryMode: amqp.Persistent,
	}

	err = ch.Publish(
		exchange,
		key,
		false,
		false,
		publishsingMsg,
	)
	if err != nil {
		p.log.Error("Error in publishing message")
	}

	p.log.Infof("Message published to exchange: %s, key: %s", exchange, key)

	return nil
}