## Описание проекта

gRPC сервис для получения приложение курса USDT. Работает с парами валют: "usdtrub", "usdtusd", "usdteur".

- gRPC server: localhost:8080 (/exchange_v1.ExchangeGRPC/GetRates)
- gRPC healthcheck: localhost:8080 (/grpc.health.v1.Health/Check)
- Metrics - Prometheus UI: http://localhost:9090
- Tracing - Jaeger UI: http://localhost:16686


## Конфигурация проекта

Конфигурация осуществляется с помощью переменных окружения в файле '.env' в репозитории проекта.
Есть возможность задать параметры подключения к PostgreSQL через флаги запуска:

  `go run ./cmd/main.go -name=<db_name> -port=<db_port> -host=<db_host> -user=<db_username> -password=<db_password>`

В таком случае, параметры конфигурации БД в файле '.env' будут проигнорированы.

## Запуск проекта

  1. `git clone https://github.com/Onewaytodice/Garantex_grpc.git`
  2. `cd Garantex_grpc`
  3. `docker compose up -d`

## Makefile:

- `make build` - для сборки приложения;
- `make test` - для запуска unit-тестов;
- `make docker-build` - для сборки Docker-образа с приложением;
- `make run` - для запуска приложения;
- `make lint` - для запуска линтера golangci-lint.

## Требования:

Go version 1.23.0
