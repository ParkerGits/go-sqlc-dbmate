package controller_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ParkerGits/go-db-starter/controller"
	"github.com/ParkerGits/go-db-starter/mocks"
	"github.com/ParkerGits/go-db-starter/query"
	"github.com/ParkerGits/go-db-starter/service"
	"go.uber.org/mock/gomock"
)

func TestHandleGetAllBooks(t *testing.T) {
	t.Run("responds with JSON containing all the books", func(t *testing.T) {
		// arrange
		ctrl := gomock.NewController(t)
		bookRepository := mocks.NewMockBookRepository(ctrl)
		bookRepository.EXPECT().GetAllBooks(gomock.Any()).Return([]query.Book{
			{ID: 1, Title: "The Pragmatic Programmer", Author: "David Thomas"},
			{ID: 2, Title: "Refactoring", Author: "Martin Fowler"},
		}, nil)
		bookService := service.NewBookService(bookRepository)
		bookController := controller.NewBookController(bookService)

		// act
		request, _ := http.NewRequest("GET", "/books", nil)
		response := httptest.NewRecorder()
		bookController.HandleGetAllBooks(response, request)

		// assert
		expected := "[{\"ID\":1,\"Title\":\"The Pragmatic Programmer\",\"Author\":\"David Thomas\"},{\"ID\":2,\"Title\":\"Refactoring\",\"Author\":\"Martin Fowler\"}]\n"
		actual := response.Body.String()

		if expected != actual {
			t.Errorf("expected %s, received %s", expected, actual)
		}
	})
}
