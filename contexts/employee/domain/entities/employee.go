package entities

import (
	"errors"
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

// Validate valida los campos requeridos y reglas de negocio para Employee
func (e *Employee) Validate() error {
	if e.PersonID == "" {
		return errors.New("personID es obligatorio")
	}
	if e.Salary <= 0 {
		return errors.New("el salario debe ser mayor a 0")
	}
	if e.ContractType == "" {
		return errors.New("contractType es obligatorio")
	}
	if len(e.ContractType) > 30 {
		return errors.New("contractType demasiado largo")
	}
	if e.Position == "" {
		return errors.New("position es obligatorio")
	}
	if len(e.Position) > 50 {
		return errors.New("position demasiado largo")
	}
	if e.Department == "" {
		return errors.New("department es obligatorio")
	}
	if len(e.Department) > 50 {
		return errors.New("department demasiado largo")
	}
	if e.WorkSchedule == "" {
		return errors.New("workSchedule es obligatorio")
	}
	if len(e.WorkSchedule) > 30 {
		return errors.New("workSchedule demasiado largo")
	}
	if e.WorkLocation == "" {
		return errors.New("workLocation es obligatorio")
	}
	if len(e.WorkLocation) > 100 {
		return errors.New("workLocation demasiado largo")
	}
	if e.BankAccount == "" {
		return errors.New("bankAccount es obligatorio")
	}
	if len(e.BankAccount) > 30 {
		return errors.New("bankAccount demasiado largo")
	}
	if e.AFP == "" {
		return errors.New("AFP es obligatorio")
	}
	if len(e.AFP) > 30 {
		return errors.New("AFP demasiado largo")
	}
	if e.EPS == "" {
		return errors.New("EPS es obligatorio")
	}
	if len(e.EPS) > 50 {
		return errors.New("EPS demasiado largo")
	}
	if e.StartDate.IsZero() {
		return errors.New("startDate es obligatorio")
	}
	if e.StartDate.After(time.Now().AddDate(0, 1, 0)) {
		return errors.New("startDate no puede ser en el futuro lejano")
	}
	if e.StartDate.Before(time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)) {
		return errors.New("startDate no puede ser antes del año 2000")
	}
	return nil
}
