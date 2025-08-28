package handler

import (
	"context"
	"log"
	"net/http"
	"time"
)

type Router struct {
	mux    *http.ServeMux
	server *http.Server
}

func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("%s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)
		next.ServeHTTP(w, r)
		log.Printf("Completed in %v", time.Since(start))
	})
}

func NewRouter() *Router {
	mux := http.NewServeMux()
	wrappedMux := loggerMiddleware(mux)

	return &Router{
		mux: mux,
		server: &http.Server{
			Addr:    ":80",
			Handler: wrappedMux,
		},
	}
}

func (r *Router) SetupRoutes() {
	// define your routes here
	// r.mux.HandleFunc("GET /generate", GetGenerate)
}

func (r *Router) Start() error {
	return r.server.ListenAndServe()
}

func (r *Router) Shutdown(ctx context.Context) error {
	return r.server.Shutdown(ctx)
}
