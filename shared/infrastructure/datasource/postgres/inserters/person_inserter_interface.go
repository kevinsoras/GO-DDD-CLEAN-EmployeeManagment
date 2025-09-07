package inserters

import (
	"context"
	"database/sql"
	"github.com/kevinsoras/employee-management/shared/domain/aggregates"
)

type PersonInserter interface {
	Insert(ctx context.Context, tx *sql.Tx, person *aggregates.PersonAggregate) error
}
