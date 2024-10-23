package main

import (
	"log/slog"
	"patterns/tee/server/config"
	"patterns/tee/server/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		slog.Error("[ERROR] read config error", "error", err.Error())
		return
	}

	app.Run(cfg)
}
