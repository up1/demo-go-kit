package main

import (
	"net/http"
	"os"
	"time"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"

	"github.com/go-kit/kit/ratelimit"
	"golang.org/x/time/rate"
)

func main() {
	c := &counterService{}

	var sep endpoint.Endpoint
	sep = makeAddEndpoint(c)

	// Rate limit
	limit := rate.NewLimiter(rate.Every(time.Minute), 1)
	sep = ratelimit.NewErroringLimiter(limit)(sep)

	addHandler := httptransport.NewServer(
		sep,
		decodeAddRequest,
		encodeResponse,
	)

	http.Handle("/add", addHandler)

	port := os.Getenv("PORT")
	http.ListenAndServe(":"+port, nil)
}
