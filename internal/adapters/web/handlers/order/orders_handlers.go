package orders

import (
	"github.com/lgustavopalmieri/go-expert-challenge-cleanarch/internal/domain/order/repository"
)

type OrdersHandler struct {
	OrderRepository repository.OrderRepositoryInterface
}

func NewOrderHandler(
	OrderRepository repository.OrderRepositoryInterface,
) *OrdersHandler {
	return &OrdersHandler{
		OrderRepository: OrderRepository,
	}
}
