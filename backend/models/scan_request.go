package models

import "database/sql"

type ScanRequest struct {
	Barcode string `json:"barcode"`
	UserID  int    `json:"userId"`
}

type ScanResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type ScanLog struct {
	Process   string     		`json:"process"`
	Component string     		`json:"component"`
	ScannedBy sql.NullString 	`json:"scannedBy"`
	ScannedAt string     		`json:"scannedAt"`
}
