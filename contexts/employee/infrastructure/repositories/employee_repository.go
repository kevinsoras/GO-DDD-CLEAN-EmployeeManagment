package repository

import (
	"context"

	"github.com/kevinsoras/employee-management/contexts/employee/domain/datasource"
	"github.com/kevinsoras/employee-management/contexts/employee/domain/entities"
	"github.com/kevinsoras/employee-management/contexts/employee/domain/repositories"
)

// EmployeeRepositoryImpl implementa EmployeeRepository usando un DataSource
// Puedes cambiar interface{} por tu entidad Employee

type EmployeeRepositoryImpl struct {
	dataSource datasource.EmployeeDataSource
}

func NewEmployeeRepositoryImpl(dataSource datasource.EmployeeDataSource) repositories.EmployeeRepository {
	return &EmployeeRepositoryImpl{dataSource: dataSource}
}

func (r *EmployeeRepositoryImpl) SaveEmployee(ctx context.Context, employee *entities.Employee) error {
	return r.dataSource.SaveEmployee(ctx, employee)
}

func (r *EmployeeRepositoryImpl) GetEmployeeByID(ctx context.Context, id string) (*entities.Employee, error) {
	return r.dataSource.GetEmployeeByID(ctx, id)
}
