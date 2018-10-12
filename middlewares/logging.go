package middlewares

import (
	"net/http"
	"log"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println(r.Method, r.RequestURI)
			next.ServeHTTP(w, r)
	})
}