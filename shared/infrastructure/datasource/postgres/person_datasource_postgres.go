package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/kevinsoras/employee-management/shared/domain/aggregates"
	"github.com/kevinsoras/employee-management/shared/domain/datasource"
	"github.com/kevinsoras/employee-management/shared/domain/value_objects"
	"github.com/kevinsoras/employee-management/shared/infrastructure/datasource/postgres/inserters"
)

type PersonDataSourcePostgres struct {
	db        *sql.DB
	inserters map[value_objects.PersonType]inserters.PersonInserter
}

func NewPersonDataSourcePostgres(db *sql.DB) datasource.PersonDataSource {
	return &PersonDataSourcePostgres{
		db: db,
		inserters: map[value_objects.PersonType]inserters.PersonInserter{
			value_objects.Natural:   inserters.NewNaturalPersonInserter(),
			value_objects.Juridical: inserters.NewJuridicPersonInserter(),
		},
	}
}

func (ds *PersonDataSourcePostgres) SavePerson(ctx context.Context, agg *aggregates.PersonAggregate) (err error) {
	tx, err := ds.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	// Use the common person inserter first
	commonInserter := inserters.NewPersonInserter()
	if err = commonInserter.Insert(ctx, tx, agg); err != nil {
		return fmt.Errorf("commonInserter.Insert: %w", err)
	}

	// Then use the specific person inserter
	specificInserter, ok := ds.inserters[agg.Person.Type]
	if !ok {
		return errors.New("no specific inserter found for person type")
	}
	if err = specificInserter.Insert(ctx, tx, agg); err != nil {
		return fmt.Errorf("specificInserter.Insert: %w", err)
	}

	return nil
}

func (ds *PersonDataSourcePostgres) GetPersonByID(ctx context.Context, id string) (*aggregates.PersonAggregate, error) {
	// Aquí va la lógica SQL para obtener una persona por ID
	return nil, nil
}
