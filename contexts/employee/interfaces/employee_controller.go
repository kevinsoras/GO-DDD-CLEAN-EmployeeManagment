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
	sharedPostgres "github.com/kevinsoras/employee-management/shared/infrastructure/datasource/postgres"
	sharedRepository "github.com/kevinsoras/employee-management/shared/infrastructure/repositories"
	"github.com/kevinsoras/employee-management/shared/utils"
)

// EmployeeController handles employee-related operations
// Injects repository and db dependencies
type EmployeeController struct {
	registerEmployeeUseCase *usecases.RegisterEmployeeUseCase
}

// NewEmployeeController creates a new controller with dependencies injected
func NewEmployeeController(db *sql.DB) *EmployeeController {
	// Data sources
	dataSource := empPostgres.NewEmployeeDataSourcePostgres(db)
	dataSourcePerson := sharedPostgres.NewPersonDataSourcePostgres(db)
	// Repositories Impl
	repo := repository.NewEmployeeRepositoryImpl(dataSource)
	repoPerson := sharedRepository.NewPersonRepositoryImpl(dataSourcePerson)
	laborService := services.NewPeruvianLaborService()
	// Use cases
	registerUC := usecases.NewRegisterEmployeeUseCase(repo, repoPerson, laborService)
	return &EmployeeController{
		registerEmployeeUseCase: registerUC,
	}
}

// HandleRegister handles the employee registration HTTP request
func (c *EmployeeController) HandleRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		return
	}
	//Validate Dto
	var dto dto.EmployeeRegistrationRequest
	if err := utils.ValidateAndBind(r, &dto); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	resp, _, err := c.registerEmployeeUseCase.Execute(r.Context(), dto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(utils.ErrorResponse(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(utils.SuccessResponse("Empleado registrado exitosamente", resp))
}
