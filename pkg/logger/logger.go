package logger

import (
	"log/slog"
	"os"
)

func LoggerInit() *slog.Logger {

	log := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	return log
}
