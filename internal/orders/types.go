package orders

import (
	"context"

	repo "github.com/jayraj/myapp/internal/adapters/postgresql/sqlc"
)

type orderItem struct {
	ProductID int64 `json:"product_id"`
	Quantity  int32 `json:"quantity"`
}

type createOrderPrams struct {
	UserID int64       `json:"user_id"`
	Items  []orderItem `json:"items"`
}

type Service interface {
	PlaceOrder(ctx context.Context, tempOrder createOrderPrams) (repo.Order, error)
}
