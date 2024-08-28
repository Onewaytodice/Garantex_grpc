build:
	go build ./cmd/main.go

test:
	go test ./... -v

docker-build:
	docker build --tag exchange:dev  .

run:
	go run -race ./cmd/main.go

lint:
	golangci-lint run

generate:
	go generate ./...
