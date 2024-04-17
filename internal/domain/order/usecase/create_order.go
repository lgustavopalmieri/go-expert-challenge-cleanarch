package usecase

import (
	"github.com/lgustavopalmieri/go-expert-challenge-cleanarch/internal/domain/order/entity"
	"github.com/lgustavopalmieri/go-expert-challenge-cleanarch/internal/domain/order/repository"
	"github.com/lgustavopalmieri/go-expert-challenge-cleanarch/pkg/events"
)

type CreateOrderUseCase struct {
	OrderRepository repository.OrderRepositoryInterface
	OrderCreated    events.EventInterface
	EventDispatcher events.EventDispatcherInterface
}

func NewCreateOrderUseCase(orderRepository repository.OrderRepositoryInterface, OrderCreated events.EventInterface,
	EventDispatcher events.EventDispatcherInterface) *CreateOrderUseCase {
	return &CreateOrderUseCase{
		OrderRepository: orderRepository,
		OrderCreated:    OrderCreated,
		EventDispatcher: EventDispatcher,
	}
}

type CreateOrderInputDTO struct {
	Price float64 `json:"price"`
	Tax   float64 `json:"tax"`
}

type CreateOrderOutputDTO struct {
	OrderID    string  `json:"order_id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
	CreatedAt  string  `json:"created_at"`
}

func (uc *CreateOrderUseCase) Execute(input CreateOrderInputDTO) (*CreateOrderOutputDTO, error) {
	order, err := entity.NewOrder(
		input.Price,
		input.Tax,
	)
	if err != nil {
		return nil, err
	}
	err = uc.OrderRepository.Save(order)
	if err != nil {
		return nil, err
	}
	dto := &CreateOrderOutputDTO{
		OrderID:    order.OrderID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
		CreatedAt:  order.CreatedAt,
	}
	uc.OrderCreated.SetPayload(dto)
	uc.EventDispatcher.Dispatch(uc.OrderCreated)

	return dto, nil
}
