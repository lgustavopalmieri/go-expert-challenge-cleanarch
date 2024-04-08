package orders

import "github.com/lgustavopalmieri/go-expert-challenge-cleanarch/internal/domain/order/repository"

type OrdersHandler struct {
	OrderRepository repository.OrderRepository
}

func NewOrderHandler(
	OrderRepository repository.OrderRepository,
) *OrdersHandler{ return &OrdersHandler{
	OrderRepository: OrderRepository,
}}