package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

const (
	writeTimeout = 15 * time.Second
	readTimeout  = 15 * time.Second
)

type server struct {
	http *http.Server
}

func MakeHttpServer(handler http.Handler) *server {
	return &server{
		http: &http.Server{
			Handler:      handler,
			WriteTimeout: writeTimeout,
			ReadTimeout:  readTimeout,
		},
	}
}

func (s *server) Start(address string) error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill)

	go func() {
		<-quit
		s.shutdown()
	}()

	s.http.Addr = address
	err := s.http.ListenAndServe()

	if err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}

func (s *server) shutdown() {
	err := s.http.Shutdown(context.Background())

	if err != nil {
		log.Fatalln(err)
	}
}
