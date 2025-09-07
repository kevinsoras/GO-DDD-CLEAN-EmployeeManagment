package entities

import (
	"errors"
	"time"

	"github.com/kevinsoras/employee-management/contexts/employee/domain/value_objects"
)

// Employee representa el agregado raíz de empleado
type Employee struct {
	id               string
	personID         string
	salary           float64
	contractType     string
	startDate        time.Time
	position         string
	workSchedule     string
	department       string
	workLocation     string
	bankAccount      string
	afp              string
	eps              string
	hasCTS           bool
	hasGratification bool
	hasVacation      bool
	benefits         value_objects.Benefits
	createdAt        time.Time
	updatedAt        time.Time
}

// --- Getters ---

func (e *Employee) ID() string {
	return e.id
}

func (e *Employee) PersonID() string {
	return e.personID
}

func (e *Employee) Salary() float64 {
	return e.salary
}

func (e *Employee) ContractType() string {
	return e.contractType
}

func (e *Employee) StartDate() time.Time {
	return e.startDate
}

func (e *Employee) Position() string {
	return e.position
}

func (e *Employee) WorkSchedule() string {
	return e.workSchedule
}

func (e *Employee) Department() string {
	return e.department
}

func (e *Employee) WorkLocation() string {
	return e.workLocation
}

func (e *Employee) BankAccount() string {
	return e.bankAccount
}

func (e *Employee) AFP() string {
	return e.afp
}

func (e *Employee) EPS() string {
	return e.eps
}

func (e *Employee) HasCTS() bool {
	return e.hasCTS
}

func (e *Employee) HasGratification() bool {
	return e.hasGratification
}

func (e *Employee) HasVacation() bool {
	return e.hasVacation
}

func (e *Employee) Benefits() value_objects.Benefits {
	return e.benefits
}

func (e *Employee) CreatedAt() time.Time {
	return e.createdAt
}

func (e *Employee) UpdatedAt() time.Time {
	return e.updatedAt
}

// --- Methods ---

// AssignBenefits asigna los beneficios calculados al empleado
func (e *Employee) AssignBenefits(benefits value_objects.Benefits) {
	e.benefits = benefits
}

// Validate valida los campos requeridos y reglas de negocio para Employee
func (e *Employee) Validate() error {
	if e.personID == "" {
		return errors.New("personID es obligatorio")
	}
	if e.salary <= 0 {
		return errors.New("el salario debe ser mayor a 0")
	}
	if e.contractType == "" {
		return errors.New("contractType es obligatorio")
	}
	if len(e.contractType) > 30 {
		return errors.New("contractType demasiado largo")
	}
	if e.position == "" {
		return errors.New("position es obligatorio")
	}
	if len(e.position) > 50 {
		return errors.New("position demasiado largo")
	}
	if e.department == "" {
		return errors.New("department es obligatorio")
	}
	if len(e.department) > 50 {
		return errors.New("department demasiado largo")
	}
	if e.workSchedule == "" {
		return errors.New("workSchedule es obligatorio")
	}
	if len(e.workSchedule) > 30 {
		return errors.New("workSchedule demasiado largo")
	}
	if e.workLocation == "" {
		return errors.New("workLocation es obligatorio")
	}
	if len(e.workLocation) > 100 {
		return errors.New("workLocation demasiado largo")
	}
	if e.bankAccount == "" {
		return errors.New("bankAccount es obligatorio")
	}
	if len(e.bankAccount) > 30 {
		return errors.New("bankAccount demasiado largo")
	}
	if e.afp == "" {
		return errors.New("AFP es obligatorio")
	}
	if len(e.afp) > 30 {
		return errors.New("AFP demasiado largo")
	}
	if e.eps == "" {
		return errors.New("EPS es obligatorio")
	}
	if len(e.eps) > 50 {
		return errors.New("EPS demasiado largo")
	}
	if e.startDate.IsZero() {
		return errors.New("startDate es obligatorio")
	}
	if e.startDate.After(time.Now().AddDate(0, 1, 0)) {
		return errors.New("startDate no puede ser en el futuro lejano")
	}
	if e.startDate.Before(time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)) {
		return errors.New("startDate no puede ser antes del año 2000")
	}
	return nil
}
