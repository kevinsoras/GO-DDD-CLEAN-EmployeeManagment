package dto

import (
	"time"

	"github.com/kevinsoras/employee-management/contexts/employee/domain/entities"
	sharedDto "github.com/kevinsoras/employee-management/shared/application/dto"
	"github.com/kevinsoras/employee-management/shared/domain/aggregates"
)

type EmployeeOutput struct {
	ID               string           `json:"id"`
	PersonID         string           `json:"personId"`
	Salary           float64          `json:"salary"`
	ContractType     string           `json:"contractType"`
	StartDate        time.Time        `json:"startDate"`
	Position         string           `json:"position"`
	WorkSchedule     string           `json:"workSchedule"`
	Department       string           `json:"department"`
	WorkLocation     string           `json:"workLocation"`
	BankAccount      string           `json:"bankAccount"`
	AFP              string           `json:"afp"`
	EPS              string           `json:"eps"`
	HasCTS           bool             `json:"hasCTS"`
	HasGratification bool             `json:"hasGratification"`
	HasVacation      bool             `json:"hasVacation"`
	Benefits         BenefitsResponse `json:"benefits"`
}

type EmployeeResponse struct {
	Employment EmployeeOutput           `json:"employment"`
	Person     sharedDto.PersonResponse `json:"person"`
}

type BenefitsResponse struct {
	CTS           float64 `json:"cts"`
	Gratification float64 `json:"gratification"`
	VacationDays  int     `json:"vacationDays"`
}

func NewEmployeeResponse(e *entities.Employee, personAgg *aggregates.PersonAggregate) EmployeeResponse {
	return EmployeeResponse{
		Employment: EmployeeOutput{
			ID:               e.ID(),
			PersonID:         e.PersonID(),
			Salary:           e.Salary(),
			ContractType:     e.ContractType(),
			StartDate:        e.StartDate(),
			Position:         e.Position(),
			WorkSchedule:     e.WorkSchedule(),
			Department:       e.Department(),
			WorkLocation:     e.WorkLocation(),
			BankAccount:      e.BankAccount(),
			AFP:              e.AFP(),
			EPS:              e.EPS(),
			HasCTS:           e.HasCTS(),
			HasGratification: e.HasGratification(),
			HasVacation:      e.HasVacation(),
			Benefits: BenefitsResponse{
				CTS:           e.Benefits().CTS(),
				Gratification: e.Benefits().Gratification(),
				VacationDays:  e.Benefits().VacationDays(),
			},
		},
		Person: sharedDto.NewPersonResponse(personAgg),
	}
}
