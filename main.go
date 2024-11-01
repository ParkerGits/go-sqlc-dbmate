package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ParkerGits/go-db-starter/query"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "db/database.sqlite3")
	if err != nil {
		panic(err)
	}

	queries := query.New(db)
	ctx := context.Background()
	_, err = queries.CreateBook(ctx, query.CreateBookParams{
		Title:  "Refactoring",
		Author: "Martin Fowler",
	})
	if err != nil {
		panic(err)
	}

	books, err := queries.GetAllBooks(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println(books)
}
