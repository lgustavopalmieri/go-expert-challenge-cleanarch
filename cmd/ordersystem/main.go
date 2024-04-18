package main

import (
	"database/sql"
	"fmt"
	"log"

	// "github.com/golang-migrate/migrate/v4"
	// _ "github.com/golang-migrate/migrate/v4/database/postgres"
	// _ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/lgustavopalmieri/go-expert-challenge-cleanarch/configs"
	"github.com/lgustavopalmieri/go-expert-challenge-cleanarch/internal/adapters/web/routes"
	"github.com/lgustavopalmieri/go-expert-challenge-cleanarch/internal/adapters/web/server"
	"github.com/lgustavopalmieri/go-expert-challenge-cleanarch/internal/domain/order/event/handler"
	postgresdb "github.com/lgustavopalmieri/go-expert-challenge-cleanarch/internal/infra/database/postgres"
	"github.com/lgustavopalmieri/go-expert-challenge-cleanarch/internal/infra/event/rabbitmq"
	"github.com/lgustavopalmieri/go-expert-challenge-cleanarch/internal/logs"
	"github.com/lgustavopalmieri/go-expert-challenge-cleanarch/pkg/events"
	"github.com/streadway/amqp"

	_ "github.com/lib/pq"
)

func main() {
	cfg, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName))
	if err != nil {
		panic(err)
	}


	postgresdb.RunPostgresMigrations(cfg.MigrationURL, cfg.DBSource)

	defer db.Close()

	rabbitMQChannel := getRabbitMQChannel()

	rabbitmq.CreateExchangeAndQueue(rabbitMQChannel, "OrderCreated", "direct", "OrderCreated", "OrderCreated")

	eventDispatcher := events.NewEventDispatcher()

	eventDispatcher.Register("OrderCreated", &handler.OrderCreatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	})

	println("connected on database", db.Ping())

	webserver := server.NewWebServer(":8086")
	routes.SetupOrdersRoutes(webserver, db, eventDispatcher)

	go webserver.Start()

	select {}
}

func getRabbitMQChannel() *amqp.Channel {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672")
	logs.FailOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	logs.FailOnError(err, "Failed to open a channel")

	log.Println("RabbitMQ connected successfully")

	return ch
}
