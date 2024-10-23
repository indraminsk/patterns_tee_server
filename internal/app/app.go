package app

import (
	"os"
	"os/signal"
	"patterns/tee/server/config"
	v1 "patterns/tee/server/internal/controller/http/v1"
	"syscall"
	httpserver "tool/http/server"
	slogger "tool/logger/slog"
)

func Run(cfg *config.Config) {
	logger := slogger.NewLogger(cfg.App.Logger.Level, cfg.App.Logger.Type)

	handler := v1.NewHandler(logger)

	srv := httpserver.New(handler.ServMux, httpserver.Addr(cfg.App.HTTP.Host, cfg.App.HTTP.Port))
	logger.Info("run service", "host", cfg.App.HTTP.Host, "port", cfg.App.HTTP.Port)

	// wait signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		logger.Info("signal.Notify:", "message", s.String())
	case err := <-srv.Notify():
		logger.Error("httpServer.Notify:", "error", err)
	}

	// shutdown
	err := srv.Shutdown()
	if err != nil {
		logger.Error("httpServer.Shutdown", "error", err)
	}

}
