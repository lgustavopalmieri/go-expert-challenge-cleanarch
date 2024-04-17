package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/lgustavopalmieri/go-expert-challenge-cleanarch/configs"
	"github.com/lgustavopalmieri/go-expert-challenge-cleanarch/internal/adapters/web/routes"
	"github.com/lgustavopalmieri/go-expert-challenge-cleanarch/internal/adapters/web/server"
	"github.com/lgustavopalmieri/go-expert-challenge-cleanarch/internal/domain/order/event/handler"
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

	runDBMigration(cfg.MigrationURL, cfg.DBSource)

	defer db.Close()

	rabbitMQChannel := getRabbitMQChannel()

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

func runDBMigration(migrationURL string, dbSource string) {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatal("cannot create new migrate instance: ", err)
	}
	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("failed to run migrate up: ", err)
	}
	log.Println("db migrate successfully")
}

func getRabbitMQChannel() *amqp.Channel {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672")
	failOnError(err, "Failed to connect to RabbitMQ")
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	log.Println("RabbitMQ connected successfully")

	err = ch.ExchangeDeclare(
		"OrderCreated", // exchange name
		"direct",       // exchange type
		true,           // durable
		false,          // deletable when not used
		false,          // exclusive (deleted when channel conection closes)
		false,          // no-wait
		nil,            // additional arguments
	)
	if err != nil {
		log.Fatalf("Error declaring exchange: %v", err)
	}

	queue, err := ch.QueueDeclare(
		"OrderCreated", // queue name
		true,           // durable
		false,          // deletable when not used
		false,          // exclusive (deleted when channel conection closes)
		false,          // no-wait
		nil,            // additional arguments
	)
	if err != nil {
		log.Fatalf("Error declaring queue: %v", err)
	}

	err = ch.QueueBind(
		queue.Name,     // queue name
		"OrderCreated", // routing key
		"OrderCreated", // exchange name
		false,          // no-wait
		nil,            // additional arguments
	)
	if err != nil {
		log.Fatalf("Error on binding queue with exchange: %v", err)
	}

	fmt.Println("Exchange, queue and bind created successfully!")
	return ch
}
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
