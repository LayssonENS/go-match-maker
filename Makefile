swagger:
	swag init -g app/main.go
	swag init -g crawling/delivery/http/crawling_handler.go
	swag init -g match/delivery/grpc/match_handler.go

mock:
	mockgen -destination=internal/domain/mock/domain_mock.go -source=internal/domain/crawling.go
	mockgen -destination=internal/domain/mock/domain_mock.go -source=internal/domain/match.go

test:
	go test ./...

build:
	docker-compose build
	docker-compose up -d

run:
	go run app/main.go
