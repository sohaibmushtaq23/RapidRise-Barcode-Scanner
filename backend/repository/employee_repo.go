package repository

import (
	"database/sql"
	"rapid-rise-backend/config"
	"rapid-rise-backend/models"
	"rapid-rise-backend/utils"
)

// Get all employees (ID + Name)
func GetEmployees() ([]models.Employee, error) {
	rows, err := config.DB.Query("SELECT EmployeeID, FirstName + ' ' + LastName as EmployeeName FROM TblEmployees WHERE FirstName IS NOT NULL AND LastName IS NOT NULL")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []models.Employee
	for rows.Next() {
		var e models.Employee
		if err := rows.Scan(&e.EmployeeID, &e.EmployeeName); err != nil {
			return nil, err
		}
		employees = append(employees, e)
	}

	return employees, nil
}

// Validate employee login
func ValidateEmployeeLogin(employeeID int, password string) (bool, error) {
	var hashedPassword string
	err := config.DB.QueryRow("SELECT PasswordHash FROM TblEmployees WHERE EmployeeID=@p1", employeeID).Scan(&hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	// Compare with bcrypt
	if !utils.CheckPasswordHash(password, hashedPassword) {
		return false, nil
	}
	return true, nil
}
