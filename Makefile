LOCAL_BIN:=$(CURDIR)/bin
LOCAL_DB_DSN:="postgres://postgres:postgres@localhost:5433/shopping_store?sslmode=disable"
LOCAL_DB_NAME:=shopping_store
GO_COVER_EXCLUDE:= "(internal/mock/**/*|*_mock.go|*_minimock.go|models_gen.go|generated.go|swagger.go|*.pb.go|.pb.*.go)"

install:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
	go get google.golang.org/grpc
	brew install protobuf
	brew install clang-format
	brew install grpcurl
	export PATH=$PATH:$(go env GOPATH)/bin

bin-deps:
	GOBIN=$(LOCAL_BIN) go install github.com/golang/mock/mockgen@v1.6.0
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.0.1

build:
	go build

image:
	docker build .

gen:
	protoc --proto_path=api/shopping_store shopping_store.proto --go-grpc_out=. --go_out=.
	protoc --proto_path=api/shopping_store_v2 shopping_store_v2.proto --go-grpc_out=. --go_out=.
	protoc --proto_path=api/shopping_store_v3 shopping_store_v3.proto --go-grpc_out=. --go_out=.

db-create-migration: NAME=$NAME
db-create-migration:
	$(LOCAL_BIN)/goose -dir db/migrations postgres "$(LOCAL_DB_DSN)" create "${NAME}" sql

db-up-local:
	$(LOCAL_BIN)/goose -dir db/migrations postgres "$(LOCAL_DB_DSN)" up

db-up:
	$(LOCAL_BIN)/goose -dir db/migrations postgres "$(DATABASE_URL)" up

db-status:
	$(LOCAL_BIN)/goose -dir db/migrations postgres "$(LOCAL_DB_DSN)" status

db-migrate:
	make db-up
	make db-gen-structure

db-migrate-down:
	$(LOCAL_BIN)/goose -dir db/migrations postgres $(LOCAL_DB_DSN) down

db-reset:
	psql -c "drop database if exists not $(LOCAL_DB_NAME) with (FORCE)"
	psql -c "create database $(LOCAL_DB_NAME)"
	make db-up

run:
	go run main.go

test:
	rm -rf cover.out
	DATABASE_URL=$(LOCAL_DB_DSN) go test -race -coverprofile=cover.out ./...

coverage:
	go tool cover -html=cover.out

mocks:
	rm -rf ./internal/pkg/mocks
	$(LOCAL_BIN)/mockgen -package=shopping_store_mock -destination=./internal/pkg/mocks/repository/shopping_store_mock.go -source=./internal/pkg/repository/shopping_store.go

ping:
	grpcurl -d '{"user_id": 2, "staff": {"staff_id": 1, "count": 3}}' -plaintext localhost:8080 shopping_store.ShoppingStore/AddStaff