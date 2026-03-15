package handlers

import (
	"encoding/json"
	"net/http"
	"rapid-rise-backend/models"
	"rapid-rise-backend/services"
)

// GET /api/employees
func GetEmployeesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	employees, err := services.GetAllEmployees()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	json.NewEncoder(w).Encode(employees)
}

// POST /api/login
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.LoginResponse{
			Success: false,
			Message: "Invalid request",
		})
		return
	}

	resp, err := services.LoginEmployee(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(resp)
}
