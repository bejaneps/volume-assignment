package main

import (
	"context"
	"log"
	"time"

	"github.com/bejaneps/volume-assignment/internal/server"
	"github.com/bejaneps/volume-assignment/internal/service/calculate"
	"github.com/bejaneps/volume-assignment/pkg/runtime"
	"github.com/caarlos0/env/v6"
)

const timeout = 5 * time.Second

type config struct {
	Port               string `env:"PORT" envDefault:"8080"`
	ServerReadTimeout  int64  `env:"SERVER_READ_TIMEOUT"`
	ServerWriteTimeout int64  `env:"SERVER_WRITE_TIMEOUT"`
}

func main() {
	cfg := &config{}
	err := env.Parse(cfg)
	if err != nil {
		log.Fatal(err)
	}

	calculateService := calculate.New()

	var opts []server.Option
	if cfg.ServerReadTimeout != 0 {
		opts = append(opts, server.WithReadTimeout(cfg.ServerReadTimeout))
	}
	if cfg.ServerWriteTimeout != 0 {
		opts = append(opts, server.WithWriteTimeout(cfg.ServerWriteTimeout))
	}

	srv := server.New(
		calculateService,
		cfg.Port,
		opts...,
	)

	runtime.RunUntilSignal(
		func() error {
			return srv.ListenAndServe()
		},
		func(ctx context.Context) error {
			return srv.Close(ctx)
		},
		timeout,
	)
}
