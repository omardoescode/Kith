run:
	go run ./cmd/server

lint:
	golangci-lint run

format:
	golangci-lint fmt
