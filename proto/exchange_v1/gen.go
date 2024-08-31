package exchange_v1

//go:generate protoc -I. -I../../../ --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative exchange.proto
