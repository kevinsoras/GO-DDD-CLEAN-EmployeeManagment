package inserters

import (
	"context"
	"github.com/kevinsoras/employee-management/shared/domain/aggregates"
	"github.com/kevinsoras/employee-management/shared/infrastructure/db"
)

type personInserter struct{}

func NewPersonInserter() PersonInserter {
	return &personInserter{}
}

func (p *personInserter) Insert(ctx context.Context, querier db.Querier, agg *aggregates.PersonAggregate) error {
	person := agg.Person
	query := `INSERT INTO persons (person_id, person_type, email, phone, address, country, created_at, updated_at)
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := querier.ExecContext(ctx, query,
		person.ID, person.Type, person.Email, person.Phone, person.Address, person.Country, person.CreatedAt, person.UpdatedAt,
	)
	return err
}
