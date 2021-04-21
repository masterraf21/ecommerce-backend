.PHONY: mock test
include .env

# run: test-api test-service
run:
	./cmds/env .env go run main.go

build:
	./cmds/env .env go build main.go

dev: 
	./cmds/env .env
	gin --appPort ${PORT} --all -i main.go

test:
	go get -u github.com/kyoh86/richgo
	./cmds/env env-test richgo test -count=1 ./... -v -cover
	go mod tidy

test-repo:
	go get -u github.com/kyoh86/richgo
	./cmds/env env-test richgo test -count=1 ./repositories/mongodb -v -cover
	go mod tidy

tidy:
	go mod tidy

download:
	go mod download

migrate-test-up:
	migrate \
  	-source file://migrations \
  	-database "mysql://test:test@tcp(localhost:3316)/test" up

migrate-test-down:
	migrate \
  	-source file://migrations \
  	-database "mysql://test:test@tcp(localhost:3316)/test" down

migrate-dev-up:
	migrate \
  	-source file://migrations \
  	-database "mysql://user:password@tcp(localhost:3326)/ordering" up

migrate-dev-down:
	migrate \
  	-source file://migrations \
  	-database "mysql://user:password@tcp(localhost:3326)/ordering" down

migrate-prod-up:
	migrate \
  	-source file://migrations \
  	-database "mysql://user:password@tcp(localhost:3336)/ordering" up

migrate-prod-down:
	migrate \
  	-source file://migrations \
  	-database "mysql://user:password@tcp(localhost:3336)/ordering" down

mock:
	@mockery --dir models --all