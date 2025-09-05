package entities

import (
	"log"
	"time"

	"github.com/google/uuid"
)

// Employee representa el agregado raíz de empleado
// Puedes ajustar los campos según tu modelo de dominio

// Benefits representa los beneficios laborales calculados para un empleado
type Benefits struct {
	CTS           float64
	Gratification float64
	VacationDays  int
}

type Employee struct {
	ID               string
	PersonID         string
	Salary           float64
	ContractType     string
	StartDate        time.Time
	Position         string
	WorkSchedule     string
	Department       string
	WorkLocation     string
	BankAccount      string
	AFP              string
	EPS              string
	HasCTS           bool
	HasGratification bool
	HasVacation      bool
	Benefits         Benefits
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

// Constructor
func NewEmployee(personID string, salary float64, contractType, position, workSchedule, department, workLocation, bankAccount, afp, eps string, startDate time.Time, hasCTS, hasGratification, hasVacation bool) *Employee {
	u7, err := uuid.NewV7()
	if err != nil {
		log.Fatalf("Error al generar ID: %v", err)
	}
	return &Employee{
		ID:               u7.String(),
		PersonID:         personID,
		Salary:           salary,
		ContractType:     contractType,
		StartDate:        startDate,
		Position:         position,
		WorkSchedule:     workSchedule,
		Department:       department,
		WorkLocation:     workLocation,
		BankAccount:      bankAccount,
		AFP:              afp,
		EPS:              eps,
		HasCTS:           hasCTS,
		HasGratification: hasGratification,
		HasVacation:      hasVacation,
		Benefits:         Benefits{},
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}
}

// AssignBenefits asigna los beneficios calculados al empleado
func (e *Employee) AssignBenefits(benefits Benefits) {
	e.Benefits = benefits
}
