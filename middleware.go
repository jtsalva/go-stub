package main

import "net/http"

func CORSOriginAndHeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", config.CorsAllowOrigin)
		w.Header().Set("Access-Control-Expose-Headers", "*")
		next.ServeHTTP(w, r)
	})
}
