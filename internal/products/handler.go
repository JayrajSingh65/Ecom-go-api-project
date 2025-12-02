package products

import (
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

	//call the service ListProducts
	// return the json in http response

	products := struct {
		Products []string `json:"products"`
	}{}
	json.Write(w, http.StatusOK, products)
}
