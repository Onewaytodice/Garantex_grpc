build:
	go build ./cmd/main.go

test:
	go test ./... -v -cover

docker-build:
	docker build --tag exchange:dev  .

run:
	go run ./cmd/main.go

lint:
	golangci-lint run ./...

generate:
	go generate ./...