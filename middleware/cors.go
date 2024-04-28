package middleware

import (
	"net/http"
	"strings"
	"time"
)

// CORS middleware
type CORS struct {
	AccessControlAllowOrigin string
	AllowMethods             []string
	AllowHeaders             []string
	AllowCredentials         bool
	ExposeHeaders            []string
	MaxAge                   time.Duration
}

func EnableCORS(config CORS) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			res.Header().Set("Access-Control-Allow-Origin", config.AccessControlAllowOrigin)
			res.Header().Set("Access-Control-Allow-Method", strings.Join(config.AllowMethods, ", "))
			res.Header().Set("Access-Control-Allow-Headers", strings.Join(config.AllowHeaders, ", "))
			res.Header().Set("Access-Control-Max-Age", config.MaxAge.Abs().String())
			res.Header().Set("Access-Control-Allow-Credentials", strings.Join(config.ExposeHeaders, ", "))
			next.ServeHTTP(res, req)
		})
	}
}
