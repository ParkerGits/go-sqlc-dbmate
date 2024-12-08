package controller

import (
	"net/http"

	"github.com/ParkerGits/go-db-starter/service"
)

type BookController interface {
	HandleGetAllBooks(w http.ResponseWriter, r *http.Request)
}

type BookControllerImpl struct {
	bookService service.BookService
}

func NewBookController(bookService service.BookService) BookController {
	return &BookControllerImpl{
		bookService: bookService,
	}
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
