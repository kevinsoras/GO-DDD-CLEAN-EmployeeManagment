package aggregates

import "github.com/kevinsoras/employee-management/shared/domain/entities"

type PersonAggregate struct {
	Person          *entities.Person
	NaturalPerson   *entities.NaturalPerson
	JuridicalPerson *entities.JuridicalPerson
}

func NewPersonAggregate(person *entities.Person, np *entities.NaturalPerson, jp *entities.JuridicalPerson) *PersonAggregate {
	return &PersonAggregate{
		Person:          person,
		NaturalPerson:   np,
		JuridicalPerson: jp,
	}
}
