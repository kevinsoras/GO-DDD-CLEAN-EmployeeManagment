package datasource

import (
	"context"

	"github.com/kevinsoras/employee-management/shared/domain/aggregates"
)

type PersonDataSource interface {
	SavePerson(ctx context.Context, person *aggregates.PersonAggregate) error
	GetPersonByID(ctx context.Context, id string) (*aggregates.PersonAggregate, error)
}
