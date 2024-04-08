package orders

import (
	"encoding/json"
	"net/http"

	"github.com/lgustavopalmieri/go-expert-challenge-cleanarch/internal/domain/order/usecase"
)

func (h *OrdersHandler) ListOrders(w http.ResponseWriter, r *http.Request) {
	orderUseCase := usecase.NewListOrdersUseCase(h.OrderRepository)
	dto, err := orderUseCase.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
