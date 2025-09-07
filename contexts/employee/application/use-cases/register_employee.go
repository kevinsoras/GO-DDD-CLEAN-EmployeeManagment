package usecases

import (
	"context"
	"fmt"

	employeedto "github.com/kevinsoras/employee-management/contexts/employee/application/dto"
	"github.com/kevinsoras/employee-management/contexts/employee/domain/entities"
	"github.com/kevinsoras/employee-management/contexts/employee/domain/repositories"
	"github.com/kevinsoras/employee-management/contexts/employee/domain/services"
	"github.com/kevinsoras/employee-management/shared/application/mappers"
	"github.com/kevinsoras/employee-management/shared/domain/factories"
	sharedRepository "github.com/kevinsoras/employee-management/shared/domain/repositories"
)

// RegisterEmployeeCommand encapsulates all the information needed to register an employee.
// This follows the Command pattern.
type RegisterEmployeeCommand struct {
	Data employeedto.EmployeeRegistrationRequest
	// Future fields like ExecutingUserID, UserRole, etc. can be added here.
}

// RegisterEmployeeUseCase orchestrates the registration of an employee.
// This is the "pure" use case, containing only business logic.
type RegisterEmployeeUseCase struct {
	employeeRepo repositories.EmployeeRepository
	personRepo   sharedRepository.PersonRepository
	laborService services.LaborService
}

// NewRegisterEmployeeUseCase creates a new RegisterEmployeeUseCase.
func NewRegisterEmployeeUseCase(employeeRepo repositories.EmployeeRepository, personRepo sharedRepository.PersonRepository, laborService services.LaborService) *RegisterEmployeeUseCase {
	return &RegisterEmployeeUseCase{
		employeeRepo: employeeRepo,
		personRepo:   personRepo,
		laborService: laborService,
	}
}

// Execute contains the core business logic for registering an employee.
func (uc *RegisterEmployeeUseCase) Execute(ctx context.Context, cmd RegisterEmployeeCommand) (employeedto.EmployeeResponse, error) {
	// 1. Create person aggregate using the factory
	personReq := cmd.Data.PersonData
	personParams := mappers.ToPersonFactoryParams(personReq)
	personAgg, err := factories.CreatePerson(personParams)
	if err != nil {
		return employeedto.EmployeeResponse{}, fmt.Errorf("error creating person: %w", err)
	}
	personID := personAgg.Person.ID

	// 2. Create Employee entity using the person ID
	e := cmd.Data.EmploymentData
	employee, err := entities.NewEmployeeBuilder(personID, e.Salary, e.ContractType, e.StartDate).
		WithJobDetails(e.Position, e.Department, e.WorkSchedule, e.WorkLocation).
		WithPayroll(e.BankAccount, e.AFP, e.EPS).
		WithBenefitFlags(e.HasCTS, e.HasGratification, e.HasVacation).
		Build()
	if err != nil {
		return employeedto.EmployeeResponse{}, fmt.Errorf("error creating employee: %w", err)
	}

	// 3. Perform domain validations using a domain service
	employmentData := services.EmploymentData{
		Salary:       e.Salary,
		ContractType: e.ContractType,
	}
	if err := uc.laborService.ValidateEmployeeRegistration(employee, employmentData); err != nil {
		return employeedto.EmployeeResponse{}, fmt.Errorf("legal validation error: %w", err)
	}

	// 4. Calculate benefits
	benefits, err := uc.laborService.CalculateBenefits(employee)
	if err != nil {
		return employeedto.EmployeeResponse{}, fmt.Errorf("error calculating benefits: %w", err)
	}
	employee.AssignBenefits(benefits)

	// 5. Persist person and employee
	if err := uc.personRepo.SavePerson(ctx, personAgg); err != nil {
		return employeedto.EmployeeResponse{}, fmt.Errorf("error saving person: %w", err)
	}
	if err := uc.employeeRepo.SaveEmployee(ctx, employee); err != nil {
		return employeedto.EmployeeResponse{}, fmt.Errorf("error saving employee: %w", err)
	}

	// 6. Map to output DTO
	return employeedto.NewEmployeeResponse(employee, personAgg), nil
}
