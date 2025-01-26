package v1

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
	"github.com/oleksandr-mazur/notification/internal/provider"
)

type ResponceString struct {
	Status  int    `json:"status"`
	Message string `json:"msg"`
}

type ResponceMap struct {
	Status  int         `json:"status"`
	Message TelegramMSG `json:"msg"`
}

type ResponceUUID struct {
	Status  int       `json:"status"`
	Message string    `json:"msg"`
	UUID    uuid.UUID `json:"uuid"`
}

type TelegramMSG struct {
	ID       uuid.UUID `json:"uuid"`
	Token    string    `json:"token"`
	ClientID int64     `json:"client_id"`
	Message  string    `json:"msg"`
}

func SendMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var tMessage TelegramMSG

	if error := json.NewDecoder(r.Body).Decode(&tMessage); error != nil {
		slog.Warn("Invalid request payload")
		json.NewEncoder(w).Encode(ResponceString{Message: "Invalid request payload", Status: 400})
		return
	}

	task_uuid := uuid.New()
	go provider.SendTelegram2(tMessage.Token, tMessage.ClientID, tMessage.Message)
	// w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(ResponceUUID{Message: "Task created", Status: 201, UUID: task_uuid})
}

func GetMessageStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}
