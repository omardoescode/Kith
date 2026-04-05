run:
	go run cmd/server.go

lint:
	golangci-lint run

format:
	golangci-lint fmt
