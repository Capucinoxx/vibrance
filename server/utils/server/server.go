package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func Start() {
	s := http.Server{Addr: ":10020"}

	go func() {
		log.Fatal(s.ListenAndServe())
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Fatal(s.Shutdown(ctx))
}
