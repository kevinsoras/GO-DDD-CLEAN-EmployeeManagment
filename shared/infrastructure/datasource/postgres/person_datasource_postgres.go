package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/kevinsoras/employee-management/shared/domain/aggregates"
	"github.com/kevinsoras/employee-management/shared/domain/datasource"
	"github.com/kevinsoras/employee-management/shared/domain/entities"
	"github.com/kevinsoras/employee-management/shared/domain/value_objects"
)

type PersonDataSourcePostgres struct {
	db *sql.DB
}

func NewPersonDataSourcePostgres(db *sql.DB) datasource.PersonDataSource {
	return &PersonDataSourcePostgres{db: db}
}

func (ds *PersonDataSourcePostgres) SavePerson(ctx context.Context, agg *aggregates.PersonAggregate) error {
	tx, err := ds.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	if err = insertPerson(ctx, tx, agg.Person); err != nil {
		return fmt.Errorf("insertPerson: %w", err)
	}

	switch agg.Person.Type {
	case value_objects.PersonType("NATURAL"):
		if agg.NaturalPerson != nil {
			if err = insertNaturalPerson(ctx, tx, agg.NaturalPerson); err != nil {
				return fmt.Errorf("insertNaturalPerson: %w", err)
			}
		}
	case value_objects.PersonType("JURIDICAL"):
		if agg.JuridicalPerson != nil {
			if err = insertJuridicalPerson(ctx, tx, agg.JuridicalPerson); err != nil {
				return fmt.Errorf("insertJuridicalPerson: %w", err)
			}
		}
	}
	return nil
}

func insertPerson(ctx context.Context, tx *sql.Tx, p *entities.Person) error {
	query := `INSERT INTO persons (person_id, person_type, email, phone, address, country, created_at, updated_at)
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := tx.ExecContext(ctx, query,
		p.ID, p.Type, p.Email, p.Phone, p.Address, p.Country, p.CreatedAt, p.UpdatedAt,
	)
	return err
}

func insertNaturalPerson(ctx context.Context, tx *sql.Tx, np *entities.NaturalPerson) error {
	query := `INSERT INTO natural_persons (person_id, document_number, first_name, last_name_paternal, last_name_maternal, birth_date, gender)
			  VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := tx.ExecContext(ctx, query,
		np.PersonID, np.DocumentNumber, np.FirstName, np.LastNamePaternal, np.LastNameMaternal, np.BirthDate, np.Gender,
	)
	return err
}

func insertJuridicalPerson(ctx context.Context, tx *sql.Tx, jp *entities.JuridicalPerson) error {
	query := `INSERT INTO juridical_persons (person_id, document_number, business_name, trade_name, constitution_date, representative_name, representative_document)
			  VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := tx.ExecContext(ctx, query,
		jp.PersonID, jp.DocumentNumber, jp.BusinessName, jp.TradeName, jp.ConstitutionDate, jp.RepresentativeName, jp.RepresentativeDocument,
	)
	return err
}

func (ds *PersonDataSourcePostgres) GetPersonByID(ctx context.Context, id string) (*aggregates.PersonAggregate, error) {
	// Aquí va la lógica SQL para obtener una persona por ID
	return nil, nil
}
