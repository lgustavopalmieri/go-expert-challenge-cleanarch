package orders

import (
	"database/sql"

	"github.com/lgustavopalmieri/go-expert-challenge-cleanarch/internal/domain/order/event"
	"github.com/lgustavopalmieri/go-expert-challenge-cleanarch/internal/infra/database/postgres/orderdb"
	"github.com/lgustavopalmieri/go-expert-challenge-cleanarch/pkg/events"
	_ "github.com/lib/pq"
)

func NewWebOrderHandler(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *OrdersHandler {
	orderRepository := orderdb.NewOrderRepositoryDb(db)
	orderCreated := event.NewOrderCreated()
	webOrderHandler := NewOrderHandler(orderRepository, eventDispatcher, orderCreated)
	return webOrderHandler
}
