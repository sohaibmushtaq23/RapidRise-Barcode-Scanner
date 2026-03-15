package services

import (
	"rapid-rise-backend/models"
	"rapid-rise-backend/repository"
	"sync"
	"time"
)

var (
	lastScan = make(map[string]time.Time)
	mu       sync.Mutex
)

func IsDuplicate(barcode string) bool {
	mu.Lock()
	defer mu.Unlock()
	now := time.Now()
	if t, ok := lastScan[barcode]; ok && now.Sub(t) < 5*time.Second {
		return true
	}
	lastScan[barcode] = now
	return false
}

func ProcessScanRequest(barcode string, userID int) (models.ScanResponse, error) {
	if IsDuplicate(barcode) {   // Note: map key is barcode only; we may want per‑user? Keep as is.
		return models.ScanResponse{Success: false, Message: "Duplicate scan detected"}, nil
	}
	success, message, err := repository.ProcessScan(barcode, userID)
	if err != nil {
		return models.ScanResponse{Success: false, Message: err.Error()}, err
	}
	return models.ScanResponse{Success: success, Message: message}, nil
}