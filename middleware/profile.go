package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func Profile() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			start := time.Now()
			next.ServeHTTP(res, req)
			elapsed := time.Since(start)
			fmt.Printf("Request %v took %v\n", req.URL.Path, elapsed)
		})
	}
}
