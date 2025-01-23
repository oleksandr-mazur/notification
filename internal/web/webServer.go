package web

import (
	"fmt"
	"net/http"
	"os"

	v1 "github.com/oleksandr-mazur/notification/internal/web/api/v1"
)

var Port string = os.Getenv("PORT")

func WebServer() {
	mux := http.NewServeMux()

	// mux.HandleFunc("GET /msg/telegram", getUsers)
	mux.HandleFunc("GET /msg/telegram/{uuid}", v1.GetMessageStatus)
	mux.HandleFunc("POST /msg/telegram", v1.SendMessage)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	fmt.Printf("Listening port %s\n", "8080")
	server.ListenAndServe()

}
