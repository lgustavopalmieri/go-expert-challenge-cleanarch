package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Order struct {
	OrderID    string
	Price      float64
	Tax        float64
	FinalPrice float64
	CreatedAt  string
}

func NewOrder(price, tax float64) (*Order, error) {
	currentTime := time.Now()
	order := &Order{
		OrderID:   uuid.New().String(),
		Price:     price,
		Tax:       tax,
		CreatedAt: currentTime.Format("2006-01-02T15:04:05"),
	}
	err := order.Validate()
	if err != nil {
		return nil, err
	}
	order.CalculateFinalPrice()
	return order, nil
}

func (o *Order) Validate() error {
	if o.OrderID == "" {
		return errors.New("invalid id")
	}
	if o.Price <= 0 {
		return errors.New("invalid price")
	}
	if o.Tax <= 0 {
		return errors.New("invalid tax")
	}
	return nil
}

func (o *Order) CalculateFinalPrice() float64 {
	o.FinalPrice = o.Price + o.Tax
	return o.FinalPrice
}
