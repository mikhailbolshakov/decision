package main

import (
	"context"
	decision "github.com/mikhailbolshakov/decision"
	"github.com/mikhailbolshakov/decision/bootstrap"
	"github.com/mikhailbolshakov/decision/kit"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// init context
	ctx := kit.NewRequestCtx().Empty().WithNewRequestId().ToContext(context.Background())

	// create a new service
	s := bootstrap.New()

	l := decision.L().Mth("main").Inf("created")

	// init service
	if err := s.Init(ctx); err != nil {
		l.E(err).St().Err("initialization")
		os.Exit(1)
	}

	l.Inf("initialized")

	// start listening
	if err := s.Start(ctx); err != nil {
		l.E(err).St().Err("listen")
		os.Exit(1)
	}

	l.Inf("listening")

	// handle app close
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	l.Inf("graceful shutdown")
	s.Close(ctx)
	os.Exit(0)
}
