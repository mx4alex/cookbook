.PHONY: run build
run: 
	go run ./cmd/app/main.go
build:
	go build -o ./build/cookbook ./cmd/app/main.go