package repository

import (
	"database/sql"
	"rapid-rise-backend/config"
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
