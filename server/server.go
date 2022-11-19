package server

import (
	"net/http"
	"time"
)

func New(mux *http.ServeMux, serverAddress string) *http.Server {
	srv := &http.Server{
		Addr:         serverAddress,
		Handler:      mux,
		TLSConfig:    nil,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	return srv
}
