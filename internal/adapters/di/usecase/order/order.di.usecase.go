package order

import (
	"database/sql"

	"github.com/lgustavopalmieri/go-expert-challenge-cleanarch/internal/domain/order/usecase"
	"github.com/lgustavopalmieri/go-expert-challenge-cleanarch/internal/infra/database/postgres/orderdb"
	_ "github.com/lib/pq"
)

func NewCreateOrderUseCase(db *sql.DB) *usecase.CreateOrderUseCase {
	orderRepository := orderdb.NewOrderRepositoryDb(db)
	createOrderUseCase := usecase.NewCreateOrderUseCase(orderRepository)
	return createOrderUseCase
}
