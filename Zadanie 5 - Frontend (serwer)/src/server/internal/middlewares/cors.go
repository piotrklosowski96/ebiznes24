package middlewares

import "net/http"

// CORSMiddleware ...
type CORSMiddleware struct {
	Handler http.Handler
}

// ServeHTTP ...
func (middleware CORSMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	middleware.Handler.ServeHTTP(writer, request)
}
