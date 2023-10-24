package main

import (
	"log/slog"
	"os"

	"github.com/jinyao-lee/kafka-logger/internal/service"
)

func main() {
	if err := service.Start(); err != nil {
		slog.With("error", err).Error("failed to start service")
		os.Exit(1)
	}

	slog.Info("service terminated")
}
