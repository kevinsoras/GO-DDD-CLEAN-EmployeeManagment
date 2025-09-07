package inserters

import (
	"context"
	"database/sql"
	"github.com/kevinsoras/employee-management/shared/domain/aggregates"
)

type personInserter struct{}

func NewPersonInserter() PersonInserter {
	return &personInserter{}
}

func (p *personInserter) Insert(ctx context.Context, tx *sql.Tx, agg *aggregates.PersonAggregate) error {
	person := agg.Person
	query := `INSERT INTO persons (person_id, person_type, email, phone, address, country, created_at, updated_at)
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := tx.ExecContext(ctx, query,
		person.ID, person.Type, person.Email, person.Phone, person.Address, person.Country, person.CreatedAt, person.UpdatedAt,
	)
	return err
}
