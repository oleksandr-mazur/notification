package v1

import (
	"encoding/json"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ResponceString{Message: http.StatusText(http.StatusOK), Status: http.StatusOK})
}
