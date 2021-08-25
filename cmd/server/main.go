package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gustavares/hub-api/server/http"
	"github.com/gustavares/hub-api/server/http/router"
)

func main() {
	ctx := context.Background()

	r := router.New(
		router.WithRoutes(http.HealthcheckRoutes()...),
	)

	httpServer := http.New(r.Router.Handler)
	httpServer.Run()

	shutdown := make(chan os.Signal, 2)
	signal.Notify(shutdown, syscall.SIGINT)

	<-shutdown
	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
