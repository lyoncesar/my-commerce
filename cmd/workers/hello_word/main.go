package main

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

func failOnError(log *zap.SugaredLoggererr error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

type Logger *zap.Logger

func InitializeLogger() Logger {
}

func main() {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	log := logger.Sugar()

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()
}
