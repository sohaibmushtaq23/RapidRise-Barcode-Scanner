package models

type Employee struct {
	EmployeeID   int    `json:"id"`
	EmployeeName string `json:"name"`
}

type LoginRequest struct {
	EmployeeID int    `json:"employeeId"`
	Password   string `json:"password"`
}

type LoginResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Token   string `json:"token,omitempty"`
}
