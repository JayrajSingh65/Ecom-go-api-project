package orders

import (
	"log"
	"net/http"

	"github.com/jayraj/myapp/internal/json"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) PlaceOrder(w http.ResponseWriter, r *http.Request) {
	var tempOrder createOrderPrams

	if err := json.Read(r, &tempOrder); err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createOrder, err := h.service.PlaceOrder(r.Context(), tempOrder)

	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	json.Write(w, http.StatusCreated, createOrder)
}
