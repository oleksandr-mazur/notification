package main

import (
	"log/slog"
	"strings"

	"github.com/oleksandr-mazur/notification/internal/web"
)

func SetLogLevel(level string) {
	levels := map[string]slog.Level{
		"DEBUG": slog.LevelDebug,
		"INFO":  slog.LevelInfo,
		"WARN":  slog.LevelWarn,
		"ERROR": slog.LevelError,
	}
	slog.SetLogLoggerLevel(levels[strings.ToUpper(level)])
}

func main() {
	SetLogLevel("INFO")
	web.WebServer()
}
