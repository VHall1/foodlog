package bootstrap

import (
	"context"
	"database/sql"
	"log"
	"net/http"
)

type service struct {
	name   string
	server *http.Server
	db     *sql.DB
}

func NewService(name string) *service {
	return &service{
		name: name,
	}
}

func (s *service) NewHttpServer(handler http.Handler) *http.Server {
	if s.server == nil {
		s.server = &http.Server{
			Addr:    ":80",
			Handler: handler,
		}
	}

	return s.server
}

func (s *service) StartHttpServer() error {
	return s.server.ListenAndServe()
}

func (s *service) ShutdownHttpServer(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func (s *service) Postgres() *sql.DB {
	if s.db == nil {
		var err error
		if s.db, err = initPostgres(); err != nil {
			log.Fatal(err)
		}
	}

	return s.db
}

func (s *service) GetName() string {
	return s.name
}
