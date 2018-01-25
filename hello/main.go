package main

import (
	"net/http"
	"os"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	c := &counterService{}

	var sep endpoint.Endpoint
	sep = makeAddEndpoint(c)

	addHandler := httptransport.NewServer(
		sep,
		decodeAddRequest,
		encodeResponse,
	)

	http.Handle("/add", addHandler)

	port := os.Getenv("PORT")
	http.ListenAndServe(":"+port, nil)
}
