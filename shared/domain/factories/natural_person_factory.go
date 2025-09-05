// domain/factories/natural_factory.go
package factories

import (
	"github.com/kevinsoras/employee-management/shared/domain/aggregates"
	"github.com/kevinsoras/employee-management/shared/domain/entities"
	"github.com/kevinsoras/employee-management/shared/domain/value_objects"
)

type NaturalPersonFactory struct{}

func (f *NaturalPersonFactory) Supports(personType value_objects.PersonType) bool {
	return personType == value_objects.Natural
}

func (f *NaturalPersonFactory) Create(params PersonFactoryParams) (*aggregates.PersonAggregate, error) {
	// Validar campos requeridos para Natural

	// Crear entidad Person base
	person := entities.NewPerson(
		params.Type,
		params.Email,
		params.Phone,
		params.Address,
		params.Country,
	)

	// Crear entidad Natural
	natural, err := entities.NewNaturalPerson(
		person.ID,
		params.DocumentNumber,
		params.FirstName,
		params.LastNamePaternal,
		params.LastNameMaternal,
		params.Gender,
		params.BirthDate,
	)
	if err != nil {
		return nil, err
	}

	// Crear Aggregate con SOLO los campos necesarios
	return aggregates.NewPersonAggregate(
		person,  // Person base
		natural, // NaturalPerson
		nil,     // JuridicalPerson (nil porque es natural)
	), nil
}
