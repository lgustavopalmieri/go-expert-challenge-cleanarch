package orders

import (
	"encoding/json"
	"net/http"

	"github.com/lgustavopalmieri/go-expert-challenge-cleanarch/internal/domain/order/usecase"
)

func (h *OrdersHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto usecase.CreateOrderInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createOrder := usecase.NewCreateOrderUseCase(h.OrderRepository)
	output, err := createOrder.Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
