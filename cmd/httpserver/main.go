package main

import (
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("Hello\n")
	fmt.Printf("Goodbye\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	logger, _ := zap.NewProduction()

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	logger.Info("Added Handlers")

	logger.Info("Starting to serve")
	http.ListenAndServe(":8080", nil)
}

func anotherFunc() {
	fmt.Printf("Hi mom!")
}

func yetAnotherFunc() {
	fmt.Printf("Prow")
}
