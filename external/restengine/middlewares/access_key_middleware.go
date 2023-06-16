package middlewares

import (
	"net/http"
	"os"
)

func AccessKeyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessKey := r.URL.Query().Get("access_key")
		if accessKey == "" {
			http.Error(w, "Missing access_key", http.StatusBadRequest)
			return
		}
		if accessKey != os.Getenv("ACCESS_KEY") {
			http.Error(w, "Invalid access_key", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
