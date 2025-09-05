package repositories

import (
	"context"

	"github.com/kevinsoras/employee-management/shared/domain/aggregates"
)

type PersonRepository interface {
	SavePerson(ctx context.Context, person *aggregates.PersonAggregate) error
	GetPersonByID(ctx context.Context, id string) (*aggregates.PersonAggregate, error)
}
