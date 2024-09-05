package main

import (
	"encoding/json"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	jsonResponse := JSONResponse{
		Error:   false,
		Message: "Hello World",
	}

	payload, err := json.Marshal(jsonResponse)
	if err != nil {
		http.Error(w, "invalid json", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(payload)

}
