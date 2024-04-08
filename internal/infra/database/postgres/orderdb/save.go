package orderdb

import "github.com/lgustavopalmieri/go-expert-challenge-cleanarch/internal/domain/order/entity"

const query = `INSERT INTO orders (
	order_id,
	price,
	tax,
	final_price
) VALUES ($1, $2, $3, $4)`

func (r *OrderRepositoyDb) Save(order *entity.Order) error {
	stmt, err := r.Db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		order.OrderID,
		order.Price,
		order.Tax,
		order.FinalPrice,
	)
	if err != nil {
		return err
	}
	return nil
}