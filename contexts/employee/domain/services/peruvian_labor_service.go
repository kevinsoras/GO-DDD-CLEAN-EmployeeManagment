package services

import (
	"fmt"
	"time"

	"github.com/kevinsoras/employee-management/contexts/employee/domain/entities"
	"github.com/kevinsoras/employee-management/contexts/employee/domain/value_objects"
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
	if employmentData.Salary < 1130 {
		return fmt.Errorf("el salario no puede ser menor al mínimo vital (S/1,130)")
	}
	if employee.ContractType() == "INDEFINIDO" {
		// Lógica de validación para contrato indefinido
		if time.Since(employee.StartDate()).Hours() < 720 { // 720 horas = 30 días
			return fmt.Errorf("para contrato indefinido, la fecha de inicio debe ser al menos 30 días antes")
		}
	}
	return nil
}

// CalculateBenefits - Cálculo de beneficios según ley peruana
func (s *PeruvianLaborService) CalculateBenefits(employee *entities.Employee) (value_objects.Benefits, error) {
	// Las variables se inicializan en su valor "cero" (0.0 para float64)
	var cts, gratification float64
	var vacationDays int

	if employee.HasCTS() {
		cts = s.calculateCTS(employee.Salary())
	}
	if employee.HasGratification() {
		gratification = s.calculateGratification(employee.Salary())
	}
	vacationDays = s.calculateVacationDays(employee.StartDate())

	return value_objects.NewBenefits(cts, gratification, vacationDays)
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
