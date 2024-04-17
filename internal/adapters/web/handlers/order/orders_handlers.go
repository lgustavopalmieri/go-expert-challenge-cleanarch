package orders

import (
	"github.com/lgustavopalmieri/go-expert-challenge-cleanarch/internal/domain/order/repository"
	"github.com/lgustavopalmieri/go-expert-challenge-cleanarch/pkg/events"
)

type OrdersHandler struct {
	OrderRepository   repository.OrderRepositoryInterface
	EventDispatcher   events.EventDispatcherInterface
	OrderCreatedEvent events.EventInterface
}

func NewOrderHandler(
	OrderRepository repository.OrderRepositoryInterface,
	EventDispatcher events.EventDispatcherInterface,
	OrderCreatedEvent events.EventInterface,
) *OrdersHandler {
	return &OrdersHandler{
		OrderRepository: OrderRepository,
		EventDispatcher:   EventDispatcher,
		OrderCreatedEvent: OrderCreatedEvent,
	}
}
