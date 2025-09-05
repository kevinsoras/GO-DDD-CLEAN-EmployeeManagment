package dto

import (
	"time"

	"github.com/kevinsoras/employee-management/shared/application/dto"
)

// EmployeeRegistrationRequest - DTO principal para registro de empleado
type EmployeeRegistrationRequest struct {
	PersonData     dto.PersonRequest `json:"person" validate:"required"`
	EmploymentData EmploymentData    `json:"employment" validate:"required"`
}

// EmploymentData - Datos laborales del empleado
type EmploymentData struct {
	Salary       float64   `json:"salary" validate:"required,min=0"`
	ContractType string    `json:"contractType" validate:"required,oneof=INDEFINIDO FIJO PRACTICANTE"`
	StartDate    time.Time `json:"startDate" validate:"required"`
	Position     string    `json:"position" validate:"required"`
	WorkSchedule string    `json:"workSchedule" validate:"required"`
	Department   string    `json:"department" validate:"required"`
	WorkLocation string    `json:"workLocation"`
	BankAccount  string    `json:"bankAccount"`
	AFP          string    `json:"afp" validate:"required"`
	EPS          string    `json:"eps" validate:"required"`
	// Campos específicos de nómina peruana
	HasCTS           bool `json:"hasCTS"`
	HasGratification bool `json:"hasGratification"`
	HasVacation      bool `json:"hasVacation"`
}
