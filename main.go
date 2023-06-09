package main

import (
	"context"
	"log"
	"net"
	"os/signal"
	"syscall"
	"time"

	"github.com/m12r/router-demo/http"
	"github.com/m12r/router-demo/internal/application"
	"github.com/m12r/router-demo/internal/mapdb"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("cannot listen: %v", err)
	}
	defer l.Close()

	app := application.NewApp(mapdb.DB{})

	s := &http.Server{
		Handler: app,
	}

	//nolint:errcheck
	go s.Serve(l)

	<-ctx.Done()
	func() {
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()

		s.SetKeepAlivesEnabled(false)
		if err := s.Shutdown(ctx); err != nil {
			log.Printf("graceful shutdown failed: %v", err)
		}
	}()
}
