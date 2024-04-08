package orders

import (
	"database/sql"

	"github.com/lgustavopalmieri/go-expert-challenge-cleanarch/internal/infra/database/postgres/orderdb"
	_ "github.com/lib/pq"
)

func NewWebOrderHandler(db *sql.DB) *OrdersHandler {
	orderRepository := orderdb.NewOrderRepositoryDb(db)
	webOrderHandler := NewOrderHandler(orderRepository)
	return webOrderHandler
}
