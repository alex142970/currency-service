currency-start:
	go run ./currency/cmd/currency/main.go --config=./currency/internal/config/config.example.yaml
gateway-start:
	go run ./gateway/cmd/gateway/main.go --config=./gateway/internal/config/config.example.yaml
cron:
	go run ./currency/cmd/cron/main.go --config=./currency/internal/config/config.example.yaml
migrate:
	go run ./currency/cmd/migrator/main.go --path=./currency/internal/migrations --config=./currency/internal/config/config.example.yaml
generate:
	protoc -I proto proto/currency/currency_service.proto --go_out=./pkg/generated --go_opt=paths=source_relative --go-grpc_out=./pkg/generated --go-grpc_opt=paths=source_relative
