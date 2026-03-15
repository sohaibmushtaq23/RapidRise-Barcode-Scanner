package models

type ScanRequest struct {
	Barcode string `json:"barcode"`
	UserID  int    `json:"userId"`
}

type ScanResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
