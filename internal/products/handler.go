package products

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

func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.service.ListProducts(r.Context())

	if err != nil {

		log.Printf("error: %v", err)

		http.Error(w, err.Error(), http.StatusInternalServerError)

		return

	}
	//call the service ListProducts
	// return the json in http response

	json.Write(w, http.StatusOK, products)
}
