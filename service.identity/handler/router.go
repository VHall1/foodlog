package handler

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"time"
)

type Router struct {
	mux    *http.ServeMux
	server *http.Server
}

var db *sql.DB

// TODO: move to a separate package
func getDB() (*sql.DB, error) {
	if db == nil {
		var err error
		// TODO: pull from config
		db, err = sql.Open("postgres", "host=postgres port=5432 user=postgres password=postgres dbname=postgres sslmode=disable")
		if err != nil {
			return nil, err
		}
	}
	return db, nil
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
	r.mux.HandleFunc("POST /login", GetLogin)
}

func (r *Router) Start() error {
	return r.server.ListenAndServe()
}

func (r *Router) Shutdown(ctx context.Context) error {
	return r.server.Shutdown(ctx)
}
