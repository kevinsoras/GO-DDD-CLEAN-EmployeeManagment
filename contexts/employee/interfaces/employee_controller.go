package interfaces

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/kevinsoras/employee-management/contexts/employee/application/dto"
	usecases "github.com/kevinsoras/employee-management/contexts/employee/application/use-cases"
	"github.com/kevinsoras/employee-management/shared/application"
	sharedDomain "github.com/kevinsoras/employee-management/shared/domain"
	"github.com/kevinsoras/employee-management/shared/utils"
)

// EmployeeController handles employee-related operations.
type EmployeeController struct {
	logger                  *slog.Logger
	registerEmployeeUseCase application.UseCase[usecases.RegisterEmployeeCommand, dto.EmployeeResponse]
}

// NewEmployeeController creates a new controller with dependencies wired up.
func NewEmployeeController(logger *slog.Logger, registerEmployeeUseCase application.UseCase[usecases.RegisterEmployeeCommand, dto.EmployeeResponse]) *EmployeeController {
	return &EmployeeController{
		logger:                  logger,
		registerEmployeeUseCase: registerEmployeeUseCase,
	}
}

// HandleRegister handles the employee registration HTTP request.
// @Summary Register a new employee
// @Description Register a new employee with personal and employment details.
// @Tags Employees
// @Accept json
// @Produce json
// @Param employee body dto.EmployeeRegistrationRequest true "Employee registration details"
// @Success 201 {object} utils.APIResponse "Employee registered successfully"
// @Failure 400 {object} utils.APIResponse "Bad request"
// @Failure 409 {object} utils.APIResponse "Conflict - Employee already exists"
// @Failure 500 {object} utils.APIResponse "Internal server error"
// @Router /employee [post]
func (c *EmployeeController) HandleRegister(w http.ResponseWriter, r *http.Request) {
	c.logger.Info("Received request to register employee")
	if r.Method != http.MethodPost {
		c.logger.Warn("Invalid method for employee registration", "method", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		return
	}

	var registrationDTO dto.EmployeeRegistrationRequest
	if err := utils.ValidateAndBind(r, &registrationDTO); err != nil {
		c.logger.Error("Failed to validate or bind request DTO", "error", err)
		utils.HandleHTTPError(w, c.logger, sharedDomain.NewInvalidInputError(err.Error(), err)) // Use NewInvalidInputError
		return
	}

	// Create the command object to pass to the use case
	cmd := usecases.RegisterEmployeeCommand{Data: registrationDTO}
	c.logger.Debug("Executing RegisterEmployeeCommand", "command", cmd)

	resp, err := c.registerEmployeeUseCase.Execute(r.Context(), cmd)
	if err != nil {
		utils.HandleHTTPError(w, c.logger, err)
		return
	}

	c.logger.Info("Successfully registered employee", "employeeID", resp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(utils.SuccessResponse("Empleado registrado exitosamente", resp))
}