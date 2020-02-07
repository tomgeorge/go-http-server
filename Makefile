build:
	go build -o bin/httpserver ./cmd/httpserver
test:
	go test -v ./cmd... ./pkg...
