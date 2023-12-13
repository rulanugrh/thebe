package middleware

import "net/http"

func CommonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "127.0.0.1:3000")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
		w.Header().Add("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS")
		w.Header().Add("Content-Type", "application/json")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
