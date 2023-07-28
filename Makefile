run-lint:
	golangci-lint run ./... -v

generate:
	go generate ./...

up-migration:
	migrate -path migrations -database "postgres://user:password@localhost:5432/shop-smart?sslmode=disable" up

down-migration:
	migrate -path migrations -database "postgres://user:password@localhost:5432/shop-smart?sslmode=disable" down

run-server:
	cd build && docker-compose up -d
	cd ../
	go run cmd/server/main.go