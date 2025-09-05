package repositories

import (
	"context"

	"github.com/kevinsoras/employee-management/shared/domain/aggregates"
	"github.com/kevinsoras/employee-management/shared/domain/datasource"
	"github.com/kevinsoras/employee-management/shared/domain/repositories"
)

type PersonRepositoryImpl struct {
	dataSource datasource.PersonDataSource
}

func NewPersonRepositoryImpl(ds datasource.PersonDataSource) repositories.PersonRepository {
	return &PersonRepositoryImpl{dataSource: ds}
}

func (r *PersonRepositoryImpl) SavePerson(ctx context.Context, person *aggregates.PersonAggregate) error {
	return r.dataSource.SavePerson(ctx, person)
}

func (r *PersonRepositoryImpl) GetPersonByID(ctx context.Context, id string) (*aggregates.PersonAggregate, error) {
	return r.dataSource.GetPersonByID(ctx, id)
}
