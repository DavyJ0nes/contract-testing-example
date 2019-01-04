package main

import (
	"log"
	"net/http"

	"github.com/davyj0nes/contract-testing-example/api/router"
)

func main() {
	r := router.New()

	log.Fatal(http.ListenAndServe("0.0.0.0:8081", r))
}
