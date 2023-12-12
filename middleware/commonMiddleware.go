package middleware

import "net/http"

func CommonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Content-Type", "*/*")
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		next.ServeHTTP(w, r)
	})
}
