package inserters

import (
	"context"
	"github.com/kevinsoras/employee-management/shared/domain/aggregates"
	"github.com/kevinsoras/employee-management/shared/infrastructure/db"
)

const insertJuridicPersonQuery = `INSERT INTO juridical_persons (person_id, document_number, business_name, trade_name, constitution_date, representative_name, representative_document)
VALUES ($1, $2, $3, $4, $5, $6, $7);`

type juridicPersonInserter struct{}

func NewJuridicPersonInserter() PersonInserter {
	return &juridicPersonInserter{}
}

func (j *juridicPersonInserter) Insert(ctx context.Context, querier db.Querier, agg *aggregates.PersonAggregate) error {
	jp := agg.JuridicalPerson
	_, err := querier.ExecContext(ctx, insertJuridicPersonQuery, 
		jp.PersonID, jp.DocumentNumber, jp.BusinessName, jp.TradeName, jp.ConstitutionDate, jp.RepresentativeName, jp.RepresentativeDocument,
	)
	return err
}
