// domain/factories/person_factory.go
package factories

import (
	"fmt"

	"github.com/kevinsoras/employee-management/shared/domain/aggregates"
	"github.com/kevinsoras/employee-management/shared/domain/value_objects"
)

type PersonFactory interface {
	Create(params PersonFactoryParams) (*aggregates.PersonAggregate, error)
	Supports(personType value_objects.PersonType) bool
}

func init() {
	// Registrar las fábricas disponibles (Default)
	RegisterFactory(value_objects.Natural, &NaturalPersonFactory{})
	RegisterFactory(value_objects.Juridical, &JuridicalPersonFactory{})
}

var factoryRegistry = make(map[value_objects.PersonType]PersonFactory)

func RegisterFactory(personType value_objects.PersonType, factory PersonFactory) {
	factoryRegistry[personType] = factory
}

func CreatePerson(params PersonFactoryParams) (*aggregates.PersonAggregate, error) {
	factory, exists := factoryRegistry[params.Type]
	if !exists {
		return nil, fmt.Errorf("no factory for person type: %s", params.Type)
	}
	return factory.Create(params)
}
