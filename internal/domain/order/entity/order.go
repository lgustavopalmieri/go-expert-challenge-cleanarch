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

func (o *Order) NewOrder(price, tax float64) (*Order, error) {
	currentTime := time.Now()
	order := &Order{
		OrderID:    uuid.New().String(),
		Price:      price,
		Tax:        tax,
		FinalPrice: o.CalculateFinalPrice(),
		CreatedAt:  currentTime.Format("0000-00-00T00:00:00"),
	}
	err := order.Validate()
	if err != nil {
		return nil, err
	}
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
