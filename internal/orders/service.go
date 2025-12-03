package orders

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	repo "github.com/jayraj/myapp/internal/adapters/postgresql/sqlc"
)

func IntToNumeric(n int32) pgtype.Numeric {
	var num pgtype.Numeric
	num.Int = big.NewInt(int64(n))
	num.Exp = 0
	num.Valid = true
	return num
}

func ToNumeric(val float64) pgtype.Numeric {
	var num pgtype.Numeric
	num.Scan(val)
	return num
}

var (
	ErrProductNotFound = errors.New("product not found")
	ErrProductNoStock  = errors.New("product has not enough stock")
)

type svc struct {
	repo *repo.Queries
	db   *pgx.Conn
}

func NewService(repo *repo.Queries, db *pgx.Conn) Service {
	return &svc{
		repo: repo,
		db:   db,
	}
}

func (s *svc) PlaceOrder(ctx context.Context, tempOrder createOrderPrams) (repo.Order, error) {

	//validate the playload
	if tempOrder.UserID == 0 {
		return repo.Order{}, fmt.Errorf("customer id is required")
	}

	if len(tempOrder.Items) == 0 {
		return repo.Order{}, fmt.Errorf("at least one iteam is required")
	}

	tx, err := s.db.Begin(ctx)
	if err != nil {
		return repo.Order{}, err
	}
	defer tx.Rollback(ctx)

	qtx := s.repo.WithTx(tx)
	//create order
	order, err := qtx.CreateOrder(ctx, repo.CreateOrderParams{
		UserID: tempOrder.UserID,
	})
	if err != nil {
		return repo.Order{}, err
	}
	//look the product if exits
	for _, item := range tempOrder.Items {
		product, err := s.repo.FindProductByID(ctx, item.ProductID)
		if err != nil {
			return repo.Order{}, ErrProductNotFound
		}

		if product.Quantity < item.Quantity {
			return repo.Order{}, ErrProductNoStock
		}
		//create order item

		_, err = qtx.CreateOrderItem(ctx, repo.CreateOrderItemParams{
			OrderID:   order.ID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     IntToNumeric(product.Price),
		})

		if err != nil {
			return repo.Order{}, err
		}
	}

	tx.Commit(ctx)

	return order, nil
}
