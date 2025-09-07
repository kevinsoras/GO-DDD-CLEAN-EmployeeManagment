package postgres

import (
	"context"
	"database/sql"

	"github.com/kevinsoras/employee-management/contexts/employee/domain/datasource"
	"github.com/kevinsoras/employee-management/contexts/employee/domain/entities"
)

// EmployeeDataSourcePostgres implementa EmployeeDataSource usando PostgreSQL

type EmployeeDataSourcePostgres struct {
	db *sql.DB
}

func NewEmployeeDataSourcePostgres(db *sql.DB) datasource.EmployeeDataSource {
	return &EmployeeDataSourcePostgres{db: db}
}

func (ds *EmployeeDataSourcePostgres) SaveEmployee(ctx context.Context, employee *entities.Employee) error {
	query := `INSERT INTO employees (
		employee_id, person_id, salary, contract_type, position, work_schedule, department, work_location, bank_account, afp, eps, start_date, has_cts, has_gratification, has_vacation, cts, gratification, vacation_days, created_at, updated_at
	) VALUES (
		$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, now(), now()
	)`
	_, err := ds.db.ExecContext(ctx, query,
		employee.ID(),
		employee.PersonID(),
		employee.Salary(),
		employee.ContractType(),
		employee.Position(),
		employee.WorkSchedule(),
		employee.Department(),
		employee.WorkLocation(),
		employee.BankAccount(),
		employee.AFP(),
		employee.EPS(),
		employee.StartDate(),
		employee.HasCTS(),
		employee.HasGratification(),
		employee.HasVacation(),
		employee.Benefits().CTS(),
		employee.Benefits().Gratification(),
		employee.Benefits().VacationDays(),
	)
	return err
}

func (ds *EmployeeDataSourcePostgres) GetEmployeeByID(ctx context.Context, id string) (*entities.Employee, error) {
	// Aquí iría la lógica SQL para buscar un empleado por ID
	return nil, nil
}
