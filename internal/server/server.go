package server

import (
	"context"
	"log"
	"net/http"
	"time"
)

// Server represents the service which provides functionality to listen
// and serve http (or any other protocol)
type Server interface {
	ListenAndServe() error
	Close(context.Context) error
}

// Service represents the service which provides functionality to find start
// and end airports from provided list of airports
type Service interface {
	// FindStartEndAirports accepts a list of sorted or unsorted airports
	// and returns a start and end airports
	FindStartEndAirports(airports [][]string) ([]string, error)
}

type handler struct {
	service Service
	srv     *http.Server
}

// New constructs http server
func New(service Service, port string, opts ...Option) Server {
	h := &handler{
		service: service,
	}

	h.initServer(port)

	// apply options
	for _, opt := range opts {
		opt(h)
	}

	return h
}

func (h *handler) initServer(port string) {
	router := http.NewServeMux()

	// POST /calculate
	router.Handle(
		`/calculate`,
		handleCalculatePost(h.service),
	)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	h.srv = server
}

// Serve listens and serves http
func (h *handler) ListenAndServe() error {
	log.Printf("starting server on %s\n", h.srv.Addr)

	return h.srv.ListenAndServe()
}

// Close shutdowns gracefully http server
func (h *handler) Close(ctx context.Context) error {
	return h.srv.Shutdown(ctx)
}

// Option is option for MFW Handler.
type Option func(*handler)

// WithReadTimeout sets request read timeout in seconds.
func WithReadTimeout(timeout int64) Option {
	return func(h *handler) {
		h.srv.ReadTimeout = time.Duration(timeout) * time.Second
	}
}

// WithWriteTimeout sets response write timeout in seconds.
func WithWriteTimeout(timeout int64) Option {
	return func(h *handler) {
		h.srv.WriteTimeout = time.Duration(timeout) * time.Second
	}
}
