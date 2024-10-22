package app

import (
	"log/slog"
	"os"
	"os/signal"
	v1 "patterns/tee/server/internal/controller/http/v1"
	slogger "patterns/tee/server/lib/logger"
	"syscall"
	"time"
	httpserver "tool/http/server"
)

type ILogger interface {
	Debug(message interface{}, args ...interface{})
	Info(message string, args ...interface{})
	Warn(message string, args ...interface{})
	Error(message interface{}, args ...interface{})
	Fatal(message interface{}, args ...interface{})
	Timing(message string, starting time.Time)
}

func Run() {
	logger := slogger.NewLogger(-4, 0)

	handler := v1.NewHandler(logger)

	srv := httpserver.New(handler.ServMux, httpserver.Addr("", "9001"))
	logger.Info("[INFO] run service", "host", "", "port", "9001")

	// wait signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		slog.Info("main - signal:", s.String())
	case err := <-srv.Notify():
		slog.Error("main - httpServer.Notify:", "error", err)
	}

	// shutdown
	err := srv.Shutdown()
	if err != nil {
		slog.Error("main - httpServer.Shutdown", err)
	}

}
