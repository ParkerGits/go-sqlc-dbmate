package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/ParkerGits/go-db-starter/controller"
	"github.com/ParkerGits/go-db-starter/query"
	"github.com/ParkerGits/go-db-starter/repository"
	"github.com/ParkerGits/go-db-starter/service"
	_ "github.com/mattn/go-sqlite3"
)

// controller-service-repository

// controller: http logic, calls service methods
// service: business logic, coordinate calls to different services and repositories
// repository: database logic, e.g. queries

func main() {
	db, err := sql.Open("sqlite3", "db/database.sqlite3")
	if err != nil {
		panic(err)
	}
	queries := query.New(db)

	bookRepository := repository.NewBookRepository(queries)
	bookService := service.NewBookService(bookRepository)
	bookController := controller.NewBookController(bookService)

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
