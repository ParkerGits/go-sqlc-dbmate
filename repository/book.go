package repository

import (
	"context"

	"github.com/ParkerGits/go-db-starter/query"
)

type BookRepository interface {
	GetAllBooks(ctx context.Context) ([]query.Book, error)
}

type BookRepositoryImpl struct {
	queries *query.Queries
}

func NewBookRepository(queries *query.Queries) BookRepository {
	return &BookRepositoryImpl{
		queries: queries,
	}
}

func (r *BookRepositoryImpl) GetAllBooks(ctx context.Context) ([]query.Book, error) {
	return r.queries.GetAllBooks(ctx)
}
