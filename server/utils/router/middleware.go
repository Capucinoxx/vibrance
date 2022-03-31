package router

import (
	"fmt"
	"net/http"
	"time"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	fn := func(rw http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(rw, r)

		fmt.Printf("[%-7s] %-15s %v\n", r.Method, r.URL.Path, time.Since(start))
	}
	return http.HandlerFunc(fn)
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
