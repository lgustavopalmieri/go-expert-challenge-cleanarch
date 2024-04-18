package service

import (
	"context"

	"github.com/lgustavopalmieri/go-expert-challenge-cleanarch/internal/adapters/grpc/order/orderpb"
	"github.com/lgustavopalmieri/go-expert-challenge-cleanarch/internal/domain/order/usecase"
)

type OrderService struct {
	orderpb.UnimplementedOrderServiceServer
	CreateOrderUseCase usecase.CreateOrderUseCase
	ListOrdersUseCase  usecase.ListOrdersUseCase
}

func NewOrderService(createOrderUseCase usecase.CreateOrderUseCase, listOrdersUseCase usecase.ListOrdersUseCase,
) *OrderService {
	return &OrderService{
		CreateOrderUseCase: createOrderUseCase,
		ListOrdersUseCase:  listOrdersUseCase,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, in *orderpb.CreateOrderRequest) (*orderpb.Order, error) {
	dto := usecase.CreateOrderInputDTO{
		Price: float64(in.Price),
		Tax:   float64(in.Tax),
	}
	output, err := s.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	return &orderpb.Order{
		OrderId:    output.OrderID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
		CreatedAt:  output.CreatedAt,
	}, nil
}

func (c *OrderService) ListOrders(ctx context.Context, input *orderpb.Blank) (*orderpb.OrderList, error) {
	orders, err := c.ListOrdersUseCase.Execute()
	if err != nil {
		return nil, err
	}
	var ordersResponse []*orderpb.Order

	for _, order := range orders {
		ordersResponse = append(ordersResponse, &orderpb.Order{
			OrderId:    order.OrderID,
			Price:      float32(order.Price),
			Tax:        float32(order.Tax),
			FinalPrice: float32(order.FinalPrice),
			CreatedAt:  order.CreatedAt,
		})
	}
	return &orderpb.OrderList{
		Orders: ordersResponse,
	}, nil
}
