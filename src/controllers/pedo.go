package controllers

import (
	"net/http"
	"github.com/rafaelcam/go-pedo-api/src/services"
	"strconv"
	"encoding/json"
)

// Retrieves the start and end parameters of the request and sends them to the service.
func GetPeDo(w http.ResponseWriter, r *http.Request) {
	start, err := strconv.Atoi(r.FormValue("start"))
	if err != nil {
		http.Error(w, "Invalid start value. Try again.", http.StatusBadRequest)
		return
	}

	end, err := strconv.Atoi(r.FormValue("end"))
	if err != nil {
		http.Error(w, "Invalid end value. Try again.", http.StatusBadRequest)
		return
	}

	result, err := services.GetPeDoByRange(start, end)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(result)
}