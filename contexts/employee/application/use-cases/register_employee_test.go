package usecases_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	employeedto "github.com/kevinsoras/employee-management/contexts/employee/application/dto"
	usecases "github.com/kevinsoras/employee-management/contexts/employee/application/use-cases"
	"github.com/kevinsoras/employee-management/contexts/employee/domain/entities"
	employee_value_objects "github.com/kevinsoras/employee-management/contexts/employee/domain/value_objects"
	"github.com/kevinsoras/employee-management/contexts/employee/domain/services"
	sharedDomain "github.com/kevinsoras/employee-management/shared/domain"
	"github.com/kevinsoras/employee-management/shared/domain/aggregates"
	shared_dto "github.com/kevinsoras/employee-management/shared/application/dto"
	sharedInfra "github.com/kevinsoras/employee-management/shared/infrastructure"
)

// MockEmployeeRepository is a mock implementation of EmployeeRepository
type MockEmployeeRepository struct {
	mock.Mock
}

func (m *MockEmployeeRepository) SaveEmployee(ctx context.Context, employee *entities.Employee) error {
	args := m.Called(ctx, employee)
	return args.Error(0)
}

func (m *MockEmployeeRepository) GetEmployeeByID(ctx context.Context, id string) (*entities.Employee, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Employee), args.Error(1)
}

// MockPersonRepository is a mock implementation of PersonRepository
type MockPersonRepository struct {
	mock.Mock
}

func (m *MockPersonRepository) SavePerson(ctx context.Context, person *aggregates.PersonAggregate) error {
	args := m.Called(ctx, person)
	return args.Error(0)
}

func (m *MockPersonRepository) GetPersonByID(ctx context.Context, id string) (*aggregates.PersonAggregate, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*aggregates.PersonAggregate), args.Error(1)
}

// MockPeruvianLaborService is a mock implementation of LaborService
type MockPeruvianLaborService struct {
	mock.Mock
}

func (m *MockPeruvianLaborService) ValidateEmployeeRegistration(employee *entities.Employee, employmentData services.EmploymentData) error {
	args := m.Called(employee, employmentData)
	return args.Error(0)
}

func (m *MockPeruvianLaborService) CalculateBenefits(employee *entities.Employee) (employee_value_objects.Benefits, error) {
	args := m.Called(employee)
	return args.Get(0).(employee_value_objects.Benefits), args.Error(1)
}

func TestRegisterEmployeeUseCase_Execute_Success(t *testing.T) {
	// Given
	mockEmployeeRepo := new(MockEmployeeRepository)
	mockPersonRepo := new(MockPersonRepository)
	mockLaborService := new(MockPeruvianLaborService)

	useCase := usecases.NewRegisterEmployeeUseCase(mockEmployeeRepo, mockPersonRepo, mockLaborService)

	req := employeedto.EmployeeRegistrationRequest{
		PersonData: shared_dto.PersonRequest{
			Type:            "NATURAL",
			FirstName:       "John",
			LastNamePaternal: "Doe",
			LastNameMaternal: "Smith",
			Email:           "john.doe@example.com",
			Phone:           "123456789",
			Address:         "123 Main St",
			Country:         "Peru",
			DocumentNumber:  "12345678",
			BirthDate:       time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
			Gender:          "M",
		},
		EmploymentData: employeedto.EmploymentData{
			Salary:       5000.0,
			ContractType: "indefinido",
			StartDate:    time.Now(),
			Position:     "Software Engineer",
			Department:   "IT",
			WorkSchedule: "full-time",
			WorkLocation: "office",
			BankAccount:  "1234567890",
			AFP:          "Integra",
			EPS:          "Rimac",
			HasCTS:         true,
			HasGratification: true,
			HasVacation:    true,
		},
	}

	ctx := context.Background()

	// Mock expectations
	mockLaborService.On("ValidateEmployeeRegistration", mock.Anything, mock.Anything).Return(nil)
	benefits, _ := employee_value_objects.NewBenefits(1000.0, 1000.0, 1000)
	mockLaborService.On("CalculateBenefits", mock.Anything).Return(benefits, nil)
	mockPersonRepo.On("SavePerson", mock.Anything, mock.Anything).Return(nil)
	mockEmployeeRepo.On("SaveEmployee", mock.Anything, mock.Anything).Return(nil)

	// When
	employeeResp, personAgg, err := useCase.Execute(ctx, req)

	// Then
	assert.NoError(t, err)
	assert.NotNil(t, employeeResp)
	assert.NotNil(t, personAgg)
	assert.Equal(t, req.PersonData.FirstName, employeeResp.Person.FirstName)
	assert.Equal(t, req.EmploymentData.Salary, employeeResp.Employment.Salary)

	mockEmployeeRepo.AssertExpectations(t)
	mockPersonRepo.AssertExpectations(t)
	mockLaborService.AssertExpectations(t)
}

func TestRegisterEmployeeUseCase_Execute_PersonCreationError(t *testing.T) {
	// Given
	mockEmployeeRepo := new(MockEmployeeRepository)
	mockPersonRepo := new(MockPersonRepository)
	mockLaborService := new(MockPeruvianLaborService)

	useCase := usecases.NewRegisterEmployeeUseCase(mockEmployeeRepo, mockPersonRepo, mockLaborService)

	req := employeedto.EmployeeRegistrationRequest{
		PersonData: shared_dto.PersonRequest{
			Type:            "NATURAL",
			FirstName:       "", // Invalid data to trigger person creation error
			LastNamePaternal: "Doe",
			LastNameMaternal: "Smith",
			Email:           "john.doe@example.com",
			Phone:           "123456789",
			Address:         "123 Main St",
			Country:         "Peru",
			DocumentNumber:  "12345678",
			BirthDate:       time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
			Gender:          "M",
		},
		EmploymentData: employeedto.EmploymentData{
			Salary:       5000.0,
			ContractType: "indefinido",
			StartDate:    time.Now(),
			Position:     "Software Engineer",
			Department:   "IT",
			WorkSchedule: "full-time",
			WorkLocation: "office",
			BankAccount:  "1234567890",
			AFP:          "Integra",
			EPS:          "Rimac",
			HasCTS:         true,
			HasGratification: true,
			HasVacation:    true,
		},
	}

	ctx := context.Background()

	// When
	employeeResp, personAgg, err := useCase.Execute(ctx, req)

	// Then
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error creating person")
	assert.Equal(t, employeedto.EmployeeResponse{}, employeeResp)
	assert.Nil(t, personAgg)

	mockEmployeeRepo.AssertNotCalled(t, "SaveEmployee", mock.Anything, mock.Anything)
	mockPersonRepo.AssertNotCalled(t, "SavePerson", mock.Anything, mock.Anything)
	mockLaborService.AssertNotCalled(t, "ValidateEmployeeRegistration", mock.Anything, mock.Anything)
	mockLaborService.AssertNotCalled(t, "CalculateBenefits", mock.Anything)
}

func TestRegisterEmployeeUseCase_Execute_EmployeeCreationError(t *testing.T) {
	// Given
	mockEmployeeRepo := new(MockEmployeeRepository)
	mockPersonRepo := new(MockPersonRepository)
	mockLaborService := new(MockPeruvianLaborService)

	useCase := usecases.NewRegisterEmployeeUseCase(mockEmployeeRepo, mockPersonRepo, mockLaborService)

	req := employeedto.EmployeeRegistrationRequest{
		PersonData: shared_dto.PersonRequest{
			Type:            "NATURAL",
			FirstName:       "John",
			LastNamePaternal: "Doe",
			LastNameMaternal: "Smith",
			Email:           "john.doe@example.com",
			Phone:           "123456789",
			Address:         "123 Main St",
			Country:         "Peru",
			DocumentNumber:  "12345678",
			BirthDate:       time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
			Gender:          "M",
		},
		EmploymentData: employeedto.EmploymentData{
			Salary:       -100.0, // Invalid salary to trigger employee creation error
			ContractType: "indefinido",
			StartDate:    time.Now(),
			Position:     "Software Engineer",
			Department:   "IT",
			WorkSchedule: "full-time",
			WorkLocation: "office",
			BankAccount:  "1234567890",
			AFP:          "Integra",
			EPS:          "Rimac",
			HasCTS:         true,
			HasGratification: true,
			HasVacation:    true,
		},
	}

	ctx := context.Background()

	// Mock expectations for successful person creation (as employee creation happens after person creation)
	// Note: We don't mock SavePerson or SaveEmployee as they shouldn't be called
	// We also don't mock labor service as it's called after employee creation
	mockLaborService.On("ValidateEmployeeRegistration", mock.Anything, mock.Anything).Return(nil).Maybe()
	benefits, _ := employee_value_objects.NewBenefits(0.0, 0.0, 0)
	mockLaborService.On("CalculateBenefits", mock.Anything).Return(benefits, nil).Maybe()
	mockPersonRepo.On("SavePerson", mock.Anything, mock.Anything).Return(nil).Maybe()
	mockEmployeeRepo.On("SaveEmployee", mock.Anything, mock.Anything).Return(nil).Maybe()


	// When
	employeeResp, personAgg, err := useCase.Execute(ctx, req)

	// Then
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error creating employee")
	assert.Equal(t, employeedto.EmployeeResponse{}, employeeResp)
	assert.Nil(t, personAgg)

	mockEmployeeRepo.AssertNotCalled(t, "SaveEmployee", mock.Anything, mock.Anything)
	mockPersonRepo.AssertNotCalled(t, "SavePerson", mock.Anything, mock.Anything)
	mockLaborService.AssertNotCalled(t, "ValidateEmployeeRegistration", mock.Anything, mock.Anything)
	mockLaborService.AssertNotCalled(t, "CalculateBenefits", mock.Anything)
}

func TestRegisterEmployeeUseCase_Execute_LaborServiceValidationError(t *testing.T) {
	// Given
	mockEmployeeRepo := new(MockEmployeeRepository)
	mockPersonRepo := new(MockPersonRepository)
	mockLaborService := new(MockPeruvianLaborService)

	useCase := usecases.NewRegisterEmployeeUseCase(mockEmployeeRepo, mockPersonRepo, mockLaborService)

	req := employeedto.EmployeeRegistrationRequest{
		PersonData: shared_dto.PersonRequest{
			Type:            "NATURAL",
			FirstName:       "John",
			LastNamePaternal: "Doe",
			LastNameMaternal: "Smith",
			Email:           "john.doe@example.com",
			Phone:           "123456789",
			Address:         "123 Main St",
			Country:         "Peru",
			DocumentNumber:  "12345678",
			BirthDate:       time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
			Gender:          "M",
		},
		EmploymentData: employeedto.EmploymentData{
			Salary:       5000.0,
			ContractType: "indefinido",
			StartDate:    time.Now(),
			Position:     "Software Engineer",
			Department:   "IT",
			WorkSchedule: "full-time",
			WorkLocation: "office",
			BankAccount:  "1234567890",
			AFP:          "Integra",
			EPS:          "Rimac",
			HasCTS:         true,
			HasGratification: true,
			HasVacation:    true,
		},
	}

	ctx := context.Background()
	validationErr := errors.New("labor service validation failed")

	// Mock expectations
	mockLaborService.On("ValidateEmployeeRegistration", mock.Anything, mock.Anything).Return(validationErr)
	benefits, _ := employee_value_objects.NewBenefits(0.0, 0.0, 0)
	mockLaborService.On("CalculateBenefits", mock.Anything).Return(benefits, nil).Maybe() // Should not be called
	mockPersonRepo.On("SavePerson", mock.Anything, mock.Anything).Return(nil).Maybe() // Should not be called
	mockEmployeeRepo.On("SaveEmployee", mock.Anything, mock.Anything).Return(nil).Maybe() // Should not be called

	// When
	employeeResp, personAgg, err := useCase.Execute(ctx, req)

	// Then
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "legal validation error")
	assert.Equal(t, employeedto.EmployeeResponse{}, employeeResp)
	assert.Nil(t, personAgg)

	mockLaborService.AssertCalled(t, "ValidateEmployeeRegistration", mock.Anything, mock.Anything)
	mockLaborService.AssertNotCalled(t, "CalculateBenefits", mock.Anything)
	mockPersonRepo.AssertNotCalled(t, "SavePerson", mock.Anything, mock.Anything)
	mockEmployeeRepo.AssertNotCalled(t, "SaveEmployee", mock.Anything, mock.Anything)
}

func TestRegisterEmployeeUseCase_Execute_CalculateBenefitsError(t *testing.T) {
	// Given
	mockEmployeeRepo := new(MockEmployeeRepository)
	mockPersonRepo := new(MockPersonRepository)
	mockLaborService := new(MockPeruvianLaborService)

	useCase := usecases.NewRegisterEmployeeUseCase(mockEmployeeRepo, mockPersonRepo, mockLaborService)

	req := employeedto.EmployeeRegistrationRequest{
		PersonData: shared_dto.PersonRequest{
			Type:            "NATURAL",
			FirstName:       "John",
			LastNamePaternal: "Doe",
			LastNameMaternal: "Smith",
			Email:           "john.doe@example.com",
			Phone:           "123456789",
			Address:         "123 Main St",
			Country:         "Peru",
			DocumentNumber:  "12345678",
			BirthDate:       time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
			Gender:          "M",
		},
		EmploymentData: employeedto.EmploymentData{
			Salary:       5000.0,
			ContractType: "indefinido",
			StartDate:    time.Now(),
			Position:     "Software Engineer",
			Department:   "IT",
			WorkSchedule: "full-time",
			WorkLocation: "office",
			BankAccount:  "1234567890",
			AFP:          "Integra",
			EPS:          "Rimac",
			HasCTS:         true,
			HasGratification: true,
			HasVacation:    true,
		},
	}

	ctx := context.Background()
	benefitsErr := errors.New("benefits calculation failed")

	// Mock expectations
	mockLaborService.On("ValidateEmployeeRegistration", mock.Anything, mock.Anything).Return(nil)
	benefits, _ := employee_value_objects.NewBenefits(0.0, 0.0, 0)
	mockLaborService.On("CalculateBenefits", mock.Anything).Return(benefits, benefitsErr)
	mockPersonRepo.On("SavePerson", mock.Anything, mock.Anything).Return(nil).Maybe() // Should not be called
	mockEmployeeRepo.On("SaveEmployee", mock.Anything, mock.Anything).Return(nil).Maybe() // Should not be called

	// When
	employeeResp, personAgg, err := useCase.Execute(ctx, req)

	// Then
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error calculating benefits")
	assert.Equal(t, employeedto.EmployeeResponse{}, employeeResp)
	assert.Nil(t, personAgg)

	mockLaborService.AssertCalled(t, "ValidateEmployeeRegistration", mock.Anything, mock.Anything)
	mockLaborService.AssertCalled(t, "CalculateBenefits", mock.Anything)
	mockPersonRepo.AssertNotCalled(t, "SavePerson", mock.Anything, mock.Anything)
	mockEmployeeRepo.AssertNotCalled(t, "SaveEmployee", mock.Anything, mock.Anything)
}

func TestRegisterEmployeeUseCase_Execute_SavePersonError(t *testing.T) {
	// Given
	mockEmployeeRepo := new(MockEmployeeRepository)
	mockPersonRepo := new(MockPersonRepository)
	mockLaborService := new(MockPeruvianLaborService)

	useCase := usecases.NewRegisterEmployeeUseCase(mockEmployeeRepo, mockPersonRepo, mockLaborService)

	req := employeedto.EmployeeRegistrationRequest{
		PersonData: shared_dto.PersonRequest{
			Type:            "NATURAL",
			FirstName:       "John",
			LastNamePaternal: "Doe",
			LastNameMaternal: "Smith",
			Email:           "john.doe@example.com",
			Phone:           "123456789",
			Address:         "123 Main St",
			Country:         "Peru",
			DocumentNumber:  "12345678",
			BirthDate:       time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
			Gender:          "M",
		},
		EmploymentData: employeedto.EmploymentData{
			Salary:       5000.0,
			ContractType: "indefinido",
			StartDate:    time.Now(),
			Position:     "Software Engineer",
			Department:   "IT",
			WorkSchedule: "full-time",
			WorkLocation: "office",
			BankAccount:  "1234567890",
			AFP:          "Integra",
			EPS:          "Rimac",
			HasCTS:         true,
			HasGratification: true,
			HasVacation:    true,
		},
	}

	ctx := context.Background()
	savePersonErr := errors.New("failed to save person")

	// Mock expectations
	mockLaborService.On("ValidateEmployeeRegistration", mock.Anything, mock.Anything).Return(nil)
	benefits, _ := employee_value_objects.NewBenefits(0.0, 0.0, 0)
	mockLaborService.On("CalculateBenefits", mock.Anything).Return(benefits, nil)
	mockPersonRepo.On("SavePerson", mock.Anything, mock.Anything).Return(savePersonErr)
	mockEmployeeRepo.On("SaveEmployee", mock.Anything, mock.Anything).Return(nil).Maybe() // Should not be called

	// When
	employeeResp, personAgg, err := useCase.Execute(ctx, req)

	// Then
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error saving person")
	assert.Equal(t, employeedto.EmployeeResponse{}, employeeResp)
	assert.Nil(t, personAgg)

	mockLaborService.AssertCalled(t, "ValidateEmployeeRegistration", mock.Anything, mock.Anything)
	mockLaborService.AssertCalled(t, "CalculateBenefits", mock.Anything)
	mockPersonRepo.AssertCalled(t, "SavePerson", mock.Anything, mock.Anything)
	mockEmployeeRepo.AssertNotCalled(t, "SaveEmployee", mock.Anything, mock.Anything)
}

func TestRegisterEmployeeUseCase_Execute_SavePersonUniqueConstraintError(t *testing.T) {
	// Given
	mockEmployeeRepo := new(MockEmployeeRepository)
	mockPersonRepo := new(MockPersonRepository)
	mockLaborService := new(MockPeruvianLaborService)

	useCase := usecases.NewRegisterEmployeeUseCase(mockEmployeeRepo, mockPersonRepo, mockLaborService)

	req := employeedto.EmployeeRegistrationRequest{
		PersonData: shared_dto.PersonRequest{
			Type:            "NATURAL",
			FirstName:       "John",
			LastNamePaternal: "Doe",
			LastNameMaternal: "Smith",
			Email:           "john.doe@example.com",
			Phone:           "123456789",
			Address:         "123 Main St",
			Country:         "Peru",
			DocumentNumber:  "12345678",
			BirthDate:       time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
			Gender:          "M",
		},
		EmploymentData: employeedto.EmploymentData{
			Salary:       5000.0,
			ContractType: "indefinido",
			StartDate:    time.Now(),
			Position:     "Software Engineer",
			Department:   "IT",
			WorkSchedule: "full-time",
			WorkLocation: "office",
			BankAccount:  "1234567890",
			AFP:          "Integra",
			EPS:          "Rimac",
			HasCTS:         true,
			HasGratification: true,
			HasVacation:    true,
		},
	}

	ctx := context.Background()

	// Mock expectations
	mockLaborService.On("ValidateEmployeeRegistration", mock.Anything, mock.Anything).Return(nil)
	benefits, _ := employee_value_objects.NewBenefits(0.0, 0.0, 0)
	mockLaborService.On("CalculateBenefits", mock.Anything).Return(benefits, nil)
	mockPersonRepo.On("SavePerson", mock.Anything, mock.Anything).Return(sharedInfra.ErrUniqueConstraint)

	// When
	employeeResp, personAgg, err := useCase.Execute(ctx, req)

	// Then
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "person with this ID/document already exists")
	var domainErr *sharedDomain.DomainError
	assert.True(t, errors.As(err, &domainErr))
	assert.Equal(t, "ALREADY_EXISTS", domainErr.Code)
	assert.Equal(t, employeedto.EmployeeResponse{}, employeeResp)
	assert.Nil(t, personAgg)

	mockLaborService.AssertCalled(t, "ValidateEmployeeRegistration", mock.Anything, mock.Anything)
	mockLaborService.AssertCalled(t, "CalculateBenefits", mock.Anything)
	mockPersonRepo.AssertCalled(t, "SavePerson", mock.Anything, mock.Anything)
}

func TestRegisterEmployeeUseCase_Execute_SaveEmployeeError(t *testing.T) {
	// Given
	mockEmployeeRepo := new(MockEmployeeRepository)
	mockPersonRepo := new(MockPersonRepository)
	mockLaborService := new(MockPeruvianLaborService)

	useCase := usecases.NewRegisterEmployeeUseCase(mockEmployeeRepo, mockPersonRepo, mockLaborService)

	req := employeedto.EmployeeRegistrationRequest{
		PersonData: shared_dto.PersonRequest{
			Type:            "NATURAL",
			FirstName:       "John",
			LastNamePaternal: "Doe",
			LastNameMaternal: "Smith",
			Email:           "john.doe@example.com",
			Phone:           "123456789",
			Address:         "123 Main St",
			Country:         "Peru",
			DocumentNumber:  "12345678",
			BirthDate:       time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
			Gender:          "M",
		},
		EmploymentData: employeedto.EmploymentData{
			Salary:       5000.0,
			ContractType: "indefinido",
			StartDate:    time.Now(),
			Position:     "Software Engineer",
			Department:   "IT",
			WorkSchedule: "full-time",
			WorkLocation: "office",
			BankAccount:  "1234567890",
			AFP:          "Integra",
			EPS:          "Rimac",
			HasCTS:         true,
			HasGratification: true,
			HasVacation:    true,
		},
	}

	ctx := context.Background()
	saveEmployeeErr := errors.New("failed to save employee")

	// Mock expectations
	mockLaborService.On("ValidateEmployeeRegistration", mock.Anything, mock.Anything).Return(nil)
	benefits, _ := employee_value_objects.NewBenefits(0.0, 0.0, 0)
	mockLaborService.On("CalculateBenefits", mock.Anything).Return(benefits, nil)
	mockPersonRepo.On("SavePerson", mock.Anything, mock.Anything).Return(nil)
	mockEmployeeRepo.On("SaveEmployee", mock.Anything, mock.Anything).Return(saveEmployeeErr)

	// When
	employeeResp, personAgg, err := useCase.Execute(ctx, req)

	// Then
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error saving employee")
	assert.Equal(t, employeedto.EmployeeResponse{}, employeeResp)
	assert.Nil(t, personAgg)

	mockLaborService.AssertCalled(t, "ValidateEmployeeRegistration", mock.Anything, mock.Anything)
	mockLaborService.AssertCalled(t, "CalculateBenefits", mock.Anything)
	mockPersonRepo.AssertCalled(t, "SavePerson", mock.Anything, mock.Anything)
	mockEmployeeRepo.AssertCalled(t, "SaveEmployee", mock.Anything, mock.Anything)
}
