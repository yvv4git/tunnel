package application

import (
	"context"
	"errors"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

type Application interface {
	start(ctx context.Context) error
}

type application struct {
	log *slog.Logger
	app Application
}

func (a *application) Start() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if a.log == nil {
		return ErrNoLoggerProvided
	}

	if a.app == nil {
		return ErrNoApplicationProvided
	}

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		<-quit
		a.log.Info("received shutdown signal")
		cancel()
	}()

	if err := a.app.start(ctx); err != nil {
		if errors.Is(err, context.Canceled) {
			a.log.Info("application was canceled by graceful shutdown")
			return nil
		}
		return err
	}

	return nil
}
