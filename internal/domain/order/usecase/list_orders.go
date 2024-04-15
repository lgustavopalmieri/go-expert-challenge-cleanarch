package usecase

import (
	"errors"

	"github.com/lgustavopalmieri/go-expert-challenge-cleanarch/internal/domain/order/repository"
)

type ListOrdersUseCase struct {
	OrderRepository repository.OrderRepositoryInterface
}

func NewListOrdersUseCase(orderRepository repository.OrderRepositoryInterface) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: orderRepository,
	}
}

type ListOrdersOutputDTO struct {
	OrderID    string  `json:"order_id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
	CreatedAt  string  `json:"created_at"`
}

func (uc *ListOrdersUseCase) Execute() ([]ListOrdersOutputDTO, error) {
	orders, err := uc.OrderRepository.ListOrders()
	if err != nil {
		return nil, err
	}

	if len(orders) == 0 {
		return nil, errors.New("no orders to list")
	}

	var output []ListOrdersOutputDTO
	for _, order := range orders {
		output = append(output, ListOrdersOutputDTO{
			OrderID:    order.OrderID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
			CreatedAt:  order.CreatedAt,
		})
	}

	return output, nil
}
