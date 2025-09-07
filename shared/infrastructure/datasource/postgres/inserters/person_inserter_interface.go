package inserters

import (
	"context"
	"github.com/kevinsoras/employee-management/shared/domain/aggregates"
	"github.com/kevinsoras/employee-management/shared/infrastructure/db"
)

type PersonInserter interface {
	Insert(ctx context.Context, querier db.Querier, person *aggregates.PersonAggregate) error
}
