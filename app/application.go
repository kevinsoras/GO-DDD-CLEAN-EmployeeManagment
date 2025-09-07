package app

import (
	"database/sql"
	"log/slog"

	usecases "github.com/kevinsoras/employee-management/contexts/employee/application/use-cases"
	"github.com/kevinsoras/employee-management/contexts/employee/domain/services"
	empPostgres "github.com/kevinsoras/employee-management/contexts/employee/infrastructure/datasource/postgres"
	repository "github.com/kevinsoras/employee-management/contexts/employee/infrastructure/repositories"
	"github.com/kevinsoras/employee-management/contexts/employee/interfaces"
	"github.com/kevinsoras/employee-management/shared/application"
	"github.com/kevinsoras/employee-management/shared/infrastructure/db"
	sharedPostgres "github.com/kevinsoras/employee-management/shared/infrastructure/datasource/postgres"
	sharedRepository "github.com/kevinsoras/employee-management/shared/infrastructure/repositories"
)

// Application agrupa todos los componentes principales de tu aplicación.
type Application struct {
	EmployeeController *interfaces.EmployeeController
	// Aquí podrías añadir otros controladores, servicios, etc.
}

// NewApplication es la función central de ensamblaje de dependencias.
// Recibe las dependencias de nivel más bajo (DB, Logger) y construye el resto.
func NewApplication(dbConn *sql.DB, logger *slog.Logger) *Application {
	// 1. DataSources
	dataSource := empPostgres.NewEmployeeDataSourcePostgres(dbConn)
	dataSourcePerson := sharedPostgres.NewPersonDataSourcePostgres(dbConn)

	// 2. Repositorios
	repo := repository.NewEmployeeRepositoryImpl(dataSource)
	repoPerson := sharedRepository.NewPersonRepositoryImpl(dataSourcePerson)

	// 3. Servicios de Dominio
	laborService := services.NewPeruvianLaborService()

	// 4. Unit of Work
	uow := db.NewPostgresUoW(dbConn)

	// 5. Casos de Uso (puros y decorados)
	registerUC := usecases.NewRegisterEmployeeUseCase(repo, repoPerson, laborService)
	transactionalRegisterUC := application.NewTransactionalDecorator(registerUC, uow)

	// 6. Controladores (ahora con constructores más simples)
	employeeController := interfaces.NewEmployeeController(logger, transactionalRegisterUC)

	return &Application{
		EmployeeController: employeeController,
	}
}