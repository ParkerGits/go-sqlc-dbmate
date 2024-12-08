package service

import (
	"context"

	"github.com/ParkerGits/go-db-starter/query"
	"github.com/ParkerGits/go-db-starter/repository"
)

type BookService interface {
	GetAllBooks(ctx context.Context) ([]query.Book, error)
}

type BookServiceImpl struct {
	bookRepository repository.BookRepository
}

func NewBookService(bookRepository repository.BookRepository) BookService {
	return &BookServiceImpl{
		bookRepository: bookRepository,
	}
}

func (s *BookServiceImpl) GetAllBooks(ctx context.Context) ([]query.Book, error) {
	return s.bookRepository.GetAllBooks(ctx)
}
