package controllers

import (
	"net/http"
	"encoding/json"
)

type health struct {
	Status string `json:"status"`
}

// Returns application status.
func AppHealth(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(health{"OK"})
}