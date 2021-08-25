package http

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/gustavares/hub-api/server/http/router"
	"github.com/valyala/fasthttp"
)

type server struct {
	HTTPServer *fasthttp.Server
}

func New() *server {
	r := router.New(
		router.WithRoutes(HealthcheckRoutes()...),
	)

	h := r.Router.Handler

	return &server{
		HTTPServer: &fasthttp.Server{
			Handler:              h,
			ReadTimeout:          5 * time.Second,
			WriteTimeout:         10 * time.Second,
			MaxConnsPerIP:        500,
			MaxRequestsPerConn:   500,
			MaxKeepaliveDuration: 5 * time.Second,
		},
	}
}

func (s server) Run() {
	go func() {
		log.Print("Starting server")
		if err := s.HTTPServer.ListenAndServe(":8080"); err != nil {
			log.Fatal(err)
		}
	}()
}

func (s server) Shutdown(ctx context.Context) error {
	return errors.New("graceful shutdown not implemented")
}
