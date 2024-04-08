package orderdb

import "github.com/lgustavopalmieri/go-expert-challenge-cleanarch/internal/domain/order/entity"

const queryListOrders = `SELECT * FROM orders`

func (r *OrderRepositoyDb) ListOrders() ([]*entity.Order, error) {
	rows, err := r.Db.Query(queryListOrders)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []*entity.Order

	for rows.Next() {
		var order entity.Order
		if err := rows.Scan(&order.OrderID, &order.Price, &order.Tax, &order.FinalPrice, &order.CreatedAt); err != nil {
			return nil, err
		}
		orders = append(orders, &order)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return orders, nil
}
