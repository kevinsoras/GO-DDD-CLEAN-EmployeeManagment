// domain/factories/juridical_factory.go
package factories

import (
	"github.com/kevinsoras/employee-management/shared/domain/aggregates"
	"github.com/kevinsoras/employee-management/shared/domain/entities"
	"github.com/kevinsoras/employee-management/shared/domain/value_objects"
)

type JuridicalPersonFactory struct{}

func (f *JuridicalPersonFactory) Supports(personType value_objects.PersonType) bool {
	return personType == value_objects.Juridical
}

func (f *JuridicalPersonFactory) Create(params PersonFactoryParams) (*aggregates.PersonAggregate, error) {
	// Validar campos requeridos para Juridical

	// Crear entidad Person base
	person := entities.NewPerson(
		params.Type,
		params.Email,
		params.Phone,
		params.Address,
		params.Country,
	)

	// Crear entidad Juridical
	juridical, err := entities.NewJuridicalPerson(
		person.ID,
		params.DocumentNumber,
		params.BusinessName,
		params.TradeName,
		params.RepresentativeName,
		params.RepresentativeDocument,
		params.ConstitutionDate,
	)
	if err != nil {
		return nil, err
	}
	// Crear Aggregate con SOLO los campos necesarios
	return aggregates.NewPersonAggregate(
		person,    // Person base
		nil,       // NaturalPerson (nil porque es jurídica)
		juridical, // JuridicalPerson
	), nil
}
