# Craft Beautiful REST APIs with Controller-Service-Repository & Dependency Injection

This repository contains the code written as part of [this YouTube video: Craft Beautiful REST APIs with Controller-Service-Repository & Dependency Injection](https://www.youtube.com/watch?v=Fkvu_QDt5xk)

[![YouTube Video Thumbnail](https://img.youtube.com/vi/Fkvu_QDt5xk/sddefault.jpg)](https://www.youtube.com/watch?v=Fkvu_QDt5xk)

## Getting started

1. Clone this repo
2. Install [dbmate](https://github.com/amacneil/dbmate) and [sqlc](https://sqlc.dev/)
3. Run `dbmate up`
4. Run `sqlc generate`
5. You're ready to roll! Run `go run main.go` to run the program

## Creating a new migration

1. Run `dbmate new <migration-name>`
2. `dbmate` will generate a new migration file in `db/migrations`
3. Add `migrate:up` and a `migrate:down` commands
4. Run `dbmate migrate` or `dbmate up` to apply the migrations

## Creating a new query

1. Open `query.sql`
2. At the bottom of the file, add a new comment containing a name and an identifier for the result count of your query
3. Below the comment, add your SQL query
4. Run `sqlc generate` to generate the new query
