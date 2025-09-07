package services

import (
	"github.com/kevinsoras/employee-management/contexts/employee/domain/entities"
	"github.com/kevinsoras/employee-management/contexts/employee/domain/value_objects"
)

type LaborService interface {
	ValidateEmployeeRegistration(employee *entities.Employee, employmentData EmploymentData) error
	CalculateBenefits(employee *entities.Employee) (value_objects.Benefits, error)
}
