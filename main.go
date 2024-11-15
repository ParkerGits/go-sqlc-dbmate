package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ParkerGits/go-db-starter/query"
	_ "github.com/mattn/go-sqlite3"
)

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}

// controller-service-repository

// controller: http logic, calls service methods
// service: business logic, coordinate calls to different services and repositories
// repository: database logic, e.g. queries

type BookRepository interface {
	GetAllBooks(ctx context.Context) ([]query.Book, error)
}

type BookRepositoryImpl struct {
	queries *query.Queries
}

func (r *BookRepositoryImpl) GetAllBooks(ctx context.Context) ([]query.Book, error) {
	return r.queries.GetAllBooks(ctx)
}

type BookService interface {
	GetAllBooks(ctx context.Context) ([]query.Book, error)
}

type BookServiceImpl struct {
	bookRepository BookRepository
}

func (s *BookServiceImpl) GetAllBooks(ctx context.Context) ([]query.Book, error) {
	return s.bookRepository.GetAllBooks(ctx)
}

type BookController interface {
	HandleGetAllBooks(w http.ResponseWriter, r *http.Request)
}

type BookControllerImpl struct {
	bookService BookService
}

func (c *BookControllerImpl) HandleGetAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := c.bookService.GetAllBooks(r.Context())
	if err != nil {
		WriteError(w, err)
		return
	}
	if err := WriteJSON(w, http.StatusOK, books); err != nil {
		WriteError(w, err)
		return
	}
}

func main() {
	db, err := sql.Open("sqlite3", "db/database.sqlite3")
	if err != nil {
		panic(err)
	}
	queries := query.New(db)

	bookRepository := &BookRepositoryImpl{
		queries: queries,
	}
	bookService := &BookServiceImpl{
		bookRepository: bookRepository,
	}
	bookController := &BookControllerImpl{
		bookService: bookService,
	}

	router := http.NewServeMux()
	router.Handle("GET /books", http.HandlerFunc(bookController.HandleGetAllBooks))

	port := ":8080"
	server := &http.Server{
		Addr:    port,
		Handler: router,
	}

	fmt.Println("Listening on port", port)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
