package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/kevinsoras/employee-management/contexts/employee/domain/value_objects"
)

// EmployeeBuilder es el constructor para la entidad Employee.
type EmployeeBuilder struct {
	employee *Employee
}

// NewEmployeeBuilder crea una nueva instancia del builder con los campos mínimos requeridos.
func NewEmployeeBuilder(personID string, salary float64, contractType string, startDate time.Time) *EmployeeBuilder {
	return &EmployeeBuilder{
		employee: &Employee{
			personID:     personID,
			salary:       salary,
			contractType: contractType,
			startDate:    startDate,
		},
	}
}

// WithJobDetails agrupa la configuración de los detalles del puesto de trabajo.
func (b *EmployeeBuilder) WithJobDetails(position, department, workSchedule, workLocation string) *EmployeeBuilder {
	b.employee.position = position
	b.employee.department = department
	b.employee.workSchedule = workSchedule
	b.employee.workLocation = workLocation
	return b
}

// WithPayroll agrupa la configuración de la información de nómina.
func (b *EmployeeBuilder) WithPayroll(bankAccount, afp, eps string) *EmployeeBuilder {
	b.employee.bankAccount = bankAccount
	b.employee.afp = afp
	b.employee.eps = eps
	return b
}

// WithBenefitFlags agrupa la configuración de los indicadores de beneficios.
func (b *EmployeeBuilder) WithBenefitFlags(hasCTS, hasGratification, hasVacation bool) *EmployeeBuilder {
	b.employee.hasCTS = hasCTS
	b.employee.hasGratification = hasGratification
	b.employee.hasVacation = hasVacation
	return b
}

// Build finaliza la construcción, valida el objeto y lo devuelve.
func (b *EmployeeBuilder) Build() (*Employee, error) {
	u7, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	b.employee.id = u7.String()
	b.employee.createdAt = time.Now()
	b.employee.updatedAt = time.Now()

	// Inicializa con un VO de Benefits vacío. El valor real se calcula y asigna después.
	emptyBenefits, _ := value_objects.NewBenefits(0, 0, 0)
	b.employee.benefits = emptyBenefits

	if err := b.employee.Validate(); err != nil {
		return nil, err
	}

	return b.employee, nil
}
