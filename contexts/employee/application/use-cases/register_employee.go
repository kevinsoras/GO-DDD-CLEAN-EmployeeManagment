package usecases

import (
	"context"
	"errors" // Add this import
	"fmt"

	employeedto "github.com/kevinsoras/employee-management/contexts/employee/application/dto"
	"github.com/kevinsoras/employee-management/contexts/employee/domain/entities"
	"github.com/kevinsoras/employee-management/contexts/employee/domain/repositories"
	"github.com/kevinsoras/employee-management/contexts/employee/domain/services"
	"github.com/kevinsoras/employee-management/shared/application/mappers"
	sharedDomain "github.com/kevinsoras/employee-management/shared/domain"
	"github.com/kevinsoras/employee-management/shared/domain/aggregates"
	"github.com/kevinsoras/employee-management/shared/domain/factories"
	sharedRepository "github.com/kevinsoras/employee-management/shared/domain/repositories"
	sharedInfra "github.com/kevinsoras/employee-management/shared/infrastructure"
)

// RegisterEmployeeUseCase orquesta el registro de un empleado

type RegisterEmployeeUseCase struct {
	employeeRepo repositories.EmployeeRepository
	personRepo   sharedRepository.PersonRepository
	laborService *services.PeruvianLaborService
}

func NewRegisterEmployeeUseCase(employeeRepo repositories.EmployeeRepository, personRepo sharedRepository.PersonRepository, laborService *services.PeruvianLaborService) *RegisterEmployeeUseCase {
	return &RegisterEmployeeUseCase{
		employeeRepo: employeeRepo,
		personRepo:   personRepo,
		laborService: laborService,
	}
}

func (uc *RegisterEmployeeUseCase) Execute(ctx context.Context, req employeedto.EmployeeRegistrationRequest) (employeedto.EmployeeResponse, *aggregates.PersonAggregate, error) {
	// 1. Crear/agregar persona usando el factory
	personReq := req.PersonData
	personParams := mappers.ToPersonFactoryParams(personReq)
	personAgg, err := factories.CreatePerson(personParams)
	if err != nil {
		return employeedto.EmployeeResponse{}, nil, fmt.Errorf("error creating person: %w", err)
	}
	personID := personAgg.Person.ID
	// 2. Crear entidad Employee usando el ID de persona
	e := req.EmploymentData
	employee, err := entities.NewEmployeeBuilder(personID, e.Salary, e.ContractType, e.StartDate).
		WithJobDetails(e.Position, e.Department, e.WorkSchedule, e.WorkLocation).
		WithPayroll(e.BankAccount, e.AFP, e.EPS).
		WithBenefitFlags(e.HasCTS, e.HasGratification, e.HasVacation).
		Build()
	if err != nil {
		return employeedto.EmployeeResponse{}, nil, fmt.Errorf("error creating employee: %w", err)
	}

	// 2. Validaciones de dominio (servicio de dominio)
	employmentData := services.EmploymentData{
		Salary:       e.Salary,
		ContractType: e.ContractType,
	}
	if err := uc.laborService.ValidateEmployeeRegistration(employee, employmentData); err != nil {
		return employeedto.EmployeeResponse{}, nil, fmt.Errorf("legal validation error: %w", err)
	}

	// 3. Calcular beneficios
	benefits, err := uc.laborService.CalculateBenefits(employee)
	if err != nil {
		return employeedto.EmployeeResponse{}, nil, fmt.Errorf("error calculating benefits: %w", err)
	}
	employee.AssignBenefits(benefits)

	// Persistir persona y empleado
	if err := uc.personRepo.SavePerson(ctx, personAgg); err != nil {
		// Check for infrastructure unique constraint error and map to domain error
		if errors.Is(err, sharedInfra.ErrUniqueConstraint) {
			return employeedto.EmployeeResponse{}, nil, sharedDomain.NewAlreadyExistsError("person with this ID/document already exists", err)
		}
		// For any other unhandled error, return a generic error
		return employeedto.EmployeeResponse{}, nil, fmt.Errorf("error saving person : %w", err)
	}
	if err := uc.employeeRepo.SaveEmployee(ctx, employee); err != nil {
		return employeedto.EmployeeResponse{}, nil, fmt.Errorf("error saving employee: %w", err)
	}

	// Mapear a output DTO
	return employeedto.NewEmployeeResponse(employee, personAgg), personAgg, nil
}
