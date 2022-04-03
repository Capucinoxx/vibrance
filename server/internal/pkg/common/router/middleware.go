package router

import (
	"log"
	"net/http"
	"time"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

type Middlewares []Middleware

func Logger(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func(begin time.Time) {
			log.Printf("[%-7s] %q %s", r.Method, r.URL.String(), time.Since(begin))
		}(time.Now())

		h.ServeHTTP(w, r)
	})
}

func CORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Protects from MimeType Sniffing
		// w.Header().Set("X-Content-Type-Options", "nosniff")
		// Prevents browser from prefetching DNS
		w.Header().Set("X-DNS-Prefetch-Control", "off")
		// Denies website content to be served in an iframe
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("Strict-Transport-Security", "max-age=5184000; includeSubDomains")
		// Prevents Internet Explorer from executing downloads in site's context
		w.Header().Set("X-Download-Options", "noopen")
		// Minimal XSS protection
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, PATCH, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Cache-Control")
		w.Header().Set("Access-Control-Expose-Headers", "Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	}
}
