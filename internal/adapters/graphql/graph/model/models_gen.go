// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Mutation struct {
}

type Order struct {
	OrderID    string  `json:"OrderID"`
	Price      float64 `json:"Price"`
	Tax        float64 `json:"Tax"`
	FinalPrice float64 `json:"FinalPrice"`
	CreatedAt  string  `json:"CreatedAt"`
}

type OrderInput struct {
	Price float64 `json:"Price"`
	Tax   float64 `json:"Tax"`
}

type Query struct {
}
