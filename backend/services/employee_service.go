package services

import (
	"rapid-rise-backend/models"
	"rapid-rise-backend/repository"
	"rapid-rise-backend/utils"
)

// Fetch all employees
func GetAllEmployees() ([]models.Employee, error) {
	return repository.GetEmployees()
}

// Authenticate employee
func LoginEmployee(req models.LoginRequest) (models.LoginResponse, error) {
	valid, err := repository.ValidateEmployeeLogin(req.EmployeeID, req.Password)
	if err != nil {
		return models.LoginResponse{Success: false, Message: err.Error()}, err
	}
	if !valid {
		return models.LoginResponse{Success: false, Message: "Invalid ID or password"}, nil
	}
	// Generate JWT
	token, err := utils.GenerateJWT(req.EmployeeID)
	if err != nil {
		return models.LoginResponse{Success: false, Message: "Token generation failed"}, err
	}
	return models.LoginResponse{
		Success: true,
		Message: "Login successful",
		Token:   token,
	}, nil
}
