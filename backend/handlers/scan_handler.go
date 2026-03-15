package handlers

import (
	"encoding/json"
	"net/http"
	"rapid-rise-backend/models"
	"rapid-rise-backend/services"
)

func ScanHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get userID from context (set by auth middleware)
	userID, ok := r.Context().Value("userID").(int)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(models.ScanResponse{Success: false, Message: "Unauthorized"})
		return
	}

	var req models.ScanRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ScanResponse{Success: false, Message: "Invalid request"})
		return
	}

	// Use userID from token, ignore any userId in the request body
	resp, err := services.ProcessScanRequest(req.Barcode, userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(resp)
}

func GetScansHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	scans, err := services.GetTopScans()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	json.NewEncoder(w).Encode(scans)
}
