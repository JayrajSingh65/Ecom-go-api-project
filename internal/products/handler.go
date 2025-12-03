package products

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
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

func (h *handler) ProductByID(w http.ResponseWriter, r *http.Request) {
	// 1. Get id from URL param
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	// 2. Call service with ID
	product, err := h.service.ProductByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// 3. Return JSON
	w.Header().Set("Content-Type", "application/json")
	json.Write(w, http.StatusOK, product)
}
