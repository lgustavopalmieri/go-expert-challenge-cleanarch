package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"

	graphql_handler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/lgustavopalmieri/go-expert-challenge-cleanarch/configs"
	"github.com/lgustavopalmieri/go-expert-challenge-cleanarch/internal/adapters/di/usecase/order"
	"github.com/lgustavopalmieri/go-expert-challenge-cleanarch/internal/adapters/graphql/graph"
	"github.com/lgustavopalmieri/go-expert-challenge-cleanarch/internal/adapters/grpc/order/orderpb"
	"github.com/lgustavopalmieri/go-expert-challenge-cleanarch/internal/adapters/grpc/order/service"
	"github.com/lgustavopalmieri/go-expert-challenge-cleanarch/internal/adapters/web/routes"
	"github.com/lgustavopalmieri/go-expert-challenge-cleanarch/internal/adapters/web/server"
	"github.com/lgustavopalmieri/go-expert-challenge-cleanarch/internal/domain/order/event/handler"
	postgresdb "github.com/lgustavopalmieri/go-expert-challenge-cleanarch/internal/infra/database/postgres"
	"github.com/lgustavopalmieri/go-expert-challenge-cleanarch/internal/infra/event/rabbitmq"
	"github.com/lgustavopalmieri/go-expert-challenge-cleanarch/internal/logs"
	"github.com/lgustavopalmieri/go-expert-challenge-cleanarch/pkg/events"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

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

	createOrderUseCase := order.NewCreateOrderUseCase(db, eventDispatcher)
	orderListUseCase := order.NewListOrdersUseCase(db)

	grpcServer := grpc.NewServer()
	createOrderService := service.NewOrderService(*createOrderUseCase, *orderListUseCase)
	orderpb.RegisterOrderServiceServer(grpcServer, createOrderService)
	reflection.Register(grpcServer)

	fmt.Println("Starting gRPC server on port", cfg.GRPCServerPort)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.GRPCServerPort))
	if err != nil {
		panic(err)
	}
	go grpcServer.Serve(lis)

	srv := graphql_handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		CreateOrderUseCase: *createOrderUseCase,
		ListOrdersUseCase:  *orderListUseCase,
	}}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	fmt.Println("Starting GraphQL server on port", cfg.GraphQLServerPort)
	http.ListenAndServe(":"+cfg.GraphQLServerPort, nil)

	// select {}
}

func getRabbitMQChannel() *amqp.Channel {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672")
	logs.FailOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	logs.FailOnError(err, "Failed to open a channel")

	log.Println("RabbitMQ connected successfully")

	return ch
}
