package interfaces

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/kevinsoras/employee-management/contexts/employee/application/dto"
	usecases "github.com/kevinsoras/employee-management/contexts/employee/application/use-cases"
	"github.com/kevinsoras/employee-management/contexts/employee/domain/services"
	empPostgres "github.com/kevinsoras/employee-management/contexts/employee/infrastructure/datasource/postgres"
	repository "github.com/kevinsoras/employee-management/contexts/employee/infrastructure/repositories"
	"github.com/kevinsoras/employee-management/shared/application"
	sharedPostgres "github.com/kevinsoras/employee-management/shared/infrastructure/datasource/postgres"
	"github.com/kevinsoras/employee-management/shared/infrastructure/db"
	sharedRepository "github.com/kevinsoras/employee-management/shared/infrastructure/repositories"
	"github.com/kevinsoras/employee-management/shared/utils"
)

// EmployeeController handles employee-related operations.
type EmployeeController struct {
	registerEmployeeUseCase application.UseCase[usecases.RegisterEmployeeCommand, dto.EmployeeResponse]
}

// NewEmployeeController creates a new controller with dependencies wired up.
func NewEmployeeController(dbConn *sql.DB) *EmployeeController {
	// Data sources
	dataSource := empPostgres.NewEmployeeDataSourcePostgres(dbConn)
	dataSourcePerson := sharedPostgres.NewPersonDataSourcePostgres(dbConn)
	// Repositories
	repo := repository.NewEmployeeRepositoryImpl(dataSource)
	repoPerson := sharedRepository.NewPersonRepositoryImpl(dataSourcePerson)
	// Domain Services
	laborService := services.NewPeruvianLaborService()

	// Create the pure use case
	registerUC := usecases.NewRegisterEmployeeUseCase(repo, repoPerson, laborService)

	// Create the Unit of Work
	uow := db.NewPostgresUoW(dbConn)

	// Decorate the use case with transactional behavior
	transactionalRegisterUC := application.NewTransactionalDecorator(registerUC, uow)

	return &EmployeeController{
		registerEmployeeUseCase: transactionalRegisterUC,
	}
}

// HandleRegister handles the employee registration HTTP request.
func (c *EmployeeController) HandleRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		return
	}

	var registrationDTO dto.EmployeeRegistrationRequest
	if err := utils.ValidateAndBind(r, &registrationDTO); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// Create the command object to pass to the use case
	cmd := usecases.RegisterEmployeeCommand{Data: registrationDTO}

	resp, err := c.registerEmployeeUseCase.Execute(r.Context(), cmd)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(utils.ErrorResponse(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(utils.SuccessResponse("Empleado registrado exitosamente", resp))
}
