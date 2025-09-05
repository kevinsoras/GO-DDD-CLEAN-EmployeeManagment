package repositories

import (
	"context"

	"github.com/kevinsoras/employee-management/contexts/employee/domain/entities"
)

// EmployeeRepository define los métodos de persistencia para empleados
// (solo contratos, sin implementación)
type EmployeeRepository interface {
	SaveEmployee(ctx context.Context, employee *entities.Employee) error
	GetEmployeeByID(ctx context.Context, id string) (*entities.Employee, error)
}
