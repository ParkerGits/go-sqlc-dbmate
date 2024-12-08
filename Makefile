repository-mocks:
	mockgen -destination ./mocks/repository.go -package mocks github.com/ParkerGits/go-db-starter/repository BookRepository

service-mocks:
	mockgen -destination ./mocks/service.go -package mocks github.com/ParkerGits/go-db-starter/service BookService

mocks: repository-mocks service-mocks

test: mocks
	go test -v ./...
