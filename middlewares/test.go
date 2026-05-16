package middlewares

import (
	"fmt"
	"net/http"
)

func TestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("I am test middleware")
		next.ServeHTTP(w, r)
	})
}
