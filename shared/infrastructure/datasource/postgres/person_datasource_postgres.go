package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/kevinsoras/employee-management/shared/domain"
	"github.com/kevinsoras/employee-management/shared/domain/aggregates"
	"github.com/kevinsoras/employee-management/shared/domain/datasource"
	"github.com/kevinsoras/employee-management/shared/domain/value_objects"
	"github.com/kevinsoras/employee-management/shared/infrastructure"
	"github.com/kevinsoras/employee-management/shared/infrastructure/datasource/postgres/inserters"
	"github.com/kevinsoras/employee-management/shared/infrastructure/db"
	"github.com/lib/pq"
)

const uniqueViolationCode = "23505"

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
	querier := db.GetQuerier(ctx, ds.db)

	// Use the common person inserter first
	commonInserter := inserters.NewPersonInserter()
	if err = commonInserter.Insert(ctx, querier, agg); err != nil {
		return ds.handleError(err)
	}

	// Then use the specific person inserter
	specificInserter, ok := ds.inserters[agg.Person.Type]
	if !ok {
		return errors.New("no specific inserter found for person type")
	}
	if err = specificInserter.Insert(ctx, querier, agg); err != nil {
		return ds.handleError(err)
	}

	return nil
}

func (ds *PersonDataSourcePostgres) GetPersonByID(ctx context.Context, id string) (*aggregates.PersonAggregate, error) {
	// Aquí va la lógica SQL para obtener una persona por ID
	return nil, nil
}

// handleError translates specific database errors into domain errors or infrastructure errors.
func (ds *PersonDataSourcePostgres) handleError(err error) error {
	var pqErr *pq.Error
	if errors.As(err, &pqErr) {
		if pqErr.Code == uniqueViolationCode {
			return domain.NewAlreadyExistsError("La persona o el documento ya se encuentra registrado.", err)
		}
		// For any other pq.Error, wrap it as a generic DB error.
		return infrastructure.NewDBError(fmt.Sprintf("Error de base de datos: %s", pqErr.Message), err)
	}
	// For any other non-pq error, wrap it as a generic DB error.
		return infrastructure.NewDBError("Error inesperado de infraestructura", err)
}
