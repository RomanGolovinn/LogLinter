package a

import (
	"log"
	"log/slog"
)

func badLogs() {
	slog.Info("Starting server") // want "log messages must begin with a lowercase letter"

	slog.Error("ошибка подключения") // want "log messages must be in english only."

	slog.Warn("failed!!!")       // want "log messages must not contain special characters."
	slog.Debug("server dead ☠️") // want "log messages must not contain special characters."

	password := "123"
	slog.Info("user pass: " + password) // want "log messages should not contain potentially sensitive data"
	slog.Info(password)                 // want "log messages should not contain potentially sensitive data"

	slog.Error("Внимание!!!") // want "log messages must begin with a lowercase letter" "log messages must be in english only." "log messages must not contain special characters."
}

func goodLogs() {
	slog.Info("starting server")
	slog.Error("failed to connect: timeout")
	log.Fatal("shutting down")
}
