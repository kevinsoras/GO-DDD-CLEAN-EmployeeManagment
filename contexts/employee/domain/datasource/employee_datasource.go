package datasource

import (
	"context"

	"github.com/kevinsoras/employee-management/contexts/employee/domain/entities"
)

// EmployeeDataSource define el contrato para fuentes de datos de empleados
// (solo interfaz, sin implementación)
type EmployeeDataSource interface {
	SaveEmployee(ctx context.Context, employee *entities.Employee) error
	GetEmployeeByID(ctx context.Context, id string) (*entities.Employee, error)
	// Otros métodos según necesidades
}
