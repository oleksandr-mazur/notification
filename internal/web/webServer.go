package web

import (
	"log/slog"
	"net/http"
	"os"
	"strconv"

	v1 "github.com/oleksandr-mazur/notification/internal/web/api/v1"
)

// var Port string = os.Getenv("PORT")
var Port int = 8080
var Address string = "0.0.0.0"

func WebServer() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /healthcheck", v1.HealthCheck)
	mux.HandleFunc("GET /msg/telegram/{uuid}", v1.GetMessageStatus)
	mux.HandleFunc("POST /msg/telegram", v1.SendMessage)

	host := Address + ":" + strconv.Itoa(Port)

	server := &http.Server{
		Addr:    host,
		Handler: mux,
	}
	slog.Info("Starting server, Listening port " + host)
	err := server.ListenAndServe()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(98)
	}
}
