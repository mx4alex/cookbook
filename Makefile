.PHONY: run build
run: 
	go run ./cmd/app/main.go
build:
	go build -o ./build/cookbook ./cmd/app/main.go
migrate:
	migrate -path ./migrations -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' up