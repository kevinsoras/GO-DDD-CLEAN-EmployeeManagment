package services

import (
	"fmt"
	"time"

	"github.com/kevinsoras/employee-management/contexts/employee/domain/entities"
)

// PeruvianLaborService - DOMAIN SERVICE (lógica de negocio peruana)
type PeruvianLaborService struct {
	// Puede tener dependencias de otros domain services/repositorios
}

func NewPeruvianLaborService() *PeruvianLaborService {
	return &PeruvianLaborService{}
}

// ValidateEmployeeRegistration - Validaciones legales PERUANAS
type EmploymentData struct {
	Salary       float64
	ContractType string
}

func (s *PeruvianLaborService) ValidateEmployeeRegistration(
	employee *entities.Employee,
	employmentData EmploymentData, // O un DTO específico
) error {
	if employmentData.Salary < 1025 {
		return fmt.Errorf("el salario no puede ser menor al mínimo vital (S/1025)")
	}
	if employee.ContractType == "INDEFINIDO" {
		// Lógica de validación peruana
	}
	return nil
}

// CalculateBenefits - Cálculo de beneficios según ley peruana
func (s *PeruvianLaborService) CalculateBenefits(employee *entities.Employee) (entities.Benefits, error) {
	benefits := entities.Benefits{}

	if employee.HasCTS {
		benefits.CTS = s.calculateCTS(employee.Salary)
	}
	if employee.HasGratification {
		benefits.Gratification = s.calculateGratification(employee.Salary)
	}
	benefits.VacationDays = s.calculateVacationDays(employee.StartDate)
	return benefits, nil
}

// Métodos privados con fórmulas específicas peruanas
func (s *PeruvianLaborService) calculateCTS(salary float64) float64 {
	return (salary + (salary / 6)) / 12
}

func (s *PeruvianLaborService) calculateGratification(salary float64) float64 {
	return salary
}

func (s *PeruvianLaborService) calculateVacationDays(startDate time.Time) int {
	// Ejemplo: 30 días por año
	return 30
}
