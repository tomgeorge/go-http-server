build:
	go build -o bin/httpserver ./cmd/httpserver
test:
	go test -v ./cmd... ./pkg...
e2e:
	./hack/e2e_test.sh
