package repository

import "github.com/lgustavopalmieri/go-expert-challenge-cleanarch/internal/domain/order/entity"

type OrderRepository interface {
	Save(order *entity.Order) error
	ListOrders()([]*entity.Order, error)
}
