package products

import (
	"context"
	"log"

	repo "github.com/jayraj/myapp/internal/adapters/postgresql/sqlc"
)

type Service interface {
	ListProducts(ctx context.Context) ([]repo.Product, error)
	ProductByID(ctx context.Context, id int64) (repo.Product, error)
}

type svc struct {
	repo repo.Querier
}

func NewService(repo repo.Querier) Service {
	return &svc{repo: repo}
}

func (s *svc) ListProducts(ctx context.Context) ([]repo.Product, error) {
	return s.repo.ListProducts(ctx)
}

func (s *svc) ProductByID(ctx context.Context, id int64) (repo.Product, error) {
	p, err := s.repo.FindProductByID(ctx, id)
	log.Println("Repo returned:", p, err)
	return p, err // this should match your sqlc query name
}
