package order

import (
	"database/sql"

	"github.com/lgustavopalmieri/go-expert-challenge-cleanarch/internal/domain/order/event"
	"github.com/lgustavopalmieri/go-expert-challenge-cleanarch/internal/domain/order/usecase"
	"github.com/lgustavopalmieri/go-expert-challenge-cleanarch/internal/infra/database/postgres/orderdb"
	"github.com/lgustavopalmieri/go-expert-challenge-cleanarch/pkg/events"
	_ "github.com/lib/pq"
)

func NewCreateOrderUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.CreateOrderUseCase {
	orderRepository := orderdb.NewOrderRepositoryDb(db)
	orderCreated := event.NewOrderCreated()
	createOrderUseCase := usecase.NewCreateOrderUseCase(orderRepository, orderCreated, eventDispatcher)
	return createOrderUseCase
}

func NewListOrdersUseCase(db *sql.DB) *usecase.ListOrdersUseCase {
	orderRepository := orderdb.NewOrderRepositoryDb(db)
	listOrdersUseCase := usecase.NewListOrdersUseCase(orderRepository)
	return listOrdersUseCase
}
