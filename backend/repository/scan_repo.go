package repository

import (
	"database/sql"
	"rapid-rise-backend/config"
	"rapid-rise-backend/models"
)

func ProcessScan(barcode string, userID int) (bool, string, error) {
	var success bool
	var message string

	_, err := config.DB.Exec(
		"EXEC dbo.process_scan @Barcode=@barcode, @ID_User=@userid, @Success=@success OUTPUT, @Message=@message OUTPUT",
		sql.Named("barcode", barcode),
		sql.Named("userid", userID),
		sql.Named("success", sql.Out{Dest: &success}),
		sql.Named("message", sql.Out{Dest: &message}),
	)

	if err != nil {
		return false, "", err
	}

	return success, message, nil
}

func GetScanHistory() ([]models.ScanLog, error) {
	rows, err := config.DB.Query("SELECT TOP 10 Process, Component, ScannedBy, ScannedAt FROM vw_ComponentsScan ORDER BY ScanID DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var scanLog []models.ScanLog
	for rows.Next() {
		var e models.ScanLog
		if err := rows.Scan(&e.Process,&e.Component, &e.ScannedBy, &e.ScannedAt); err != nil {
			return nil, err
		}
		scanLog = append(scanLog, e)
	}

	return scanLog, nil
}