package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gustavares/hub-api/server/http"
)

func main() {
	ctx := context.Background()
	httpServer := http.New()
	httpServer.Run()

	shutdown := make(chan os.Signal, 2)
	signal.Notify(shutdown, syscall.SIGINT)

	<-shutdown
	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
