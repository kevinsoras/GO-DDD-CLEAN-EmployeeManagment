package inserters

import (
	"context"
	"github.com/kevinsoras/employee-management/shared/domain/aggregates"
	"github.com/kevinsoras/employee-management/shared/infrastructure/db"
)

const insertNaturalPersonQuery = `INSERT INTO natural_persons (person_id, document_number, first_name, last_name_paternal, last_name_maternal, birth_date, gender)
VALUES ($1, $2, $3, $4, $5, $6, $7);`

type naturalPersonInserter struct{}

func NewNaturalPersonInserter() PersonInserter {
	return &naturalPersonInserter{}
}

func (n *naturalPersonInserter) Insert(ctx context.Context, querier db.Querier, agg *aggregates.PersonAggregate) error {
	np := agg.NaturalPerson
	_, err := querier.ExecContext(ctx, insertNaturalPersonQuery, 
		np.PersonID, np.DocumentNumber, np.FirstName, np.LastNamePaternal, np.LastNameMaternal, np.BirthDate, np.Gender,
	)
	return err
}
