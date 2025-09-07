// shared/application/mappers/person_mapper.go
package mappers

import (
	"log"

	"github.com/kevinsoras/employee-management/shared/application/dto"
	"github.com/kevinsoras/employee-management/shared/domain/factories"
	"github.com/kevinsoras/employee-management/shared/domain/value_objects"
)

func ToPersonFactoryParams(personRequest dto.PersonRequest) factories.PersonFactoryParams {
	personType, err := value_objects.NewPersonType(personRequest.Type)
	if err != nil {
		log.Printf("Error creating person type: %v", err)
	}

	email, err := value_objects.NewEmail(personRequest.Email)
	if err != nil {
		log.Printf("Error creating email: %v", err)
	}

	phone, err := value_objects.NewPhone(personRequest.Phone)
	if err != nil {
		log.Printf("Error creating phone: %v", err)
	}

	params := factories.PersonFactoryParams{
		Type:           personType,
		Email:          email,
		Phone:          phone,
		Address:        personRequest.Address,
		Country:        personRequest.Country,
		DocumentNumber: personRequest.DocumentNumber,
	}
	params.FirstName = &personRequest.FirstName
	params.LastNamePaternal = &personRequest.LastNamePaternal
	params.LastNameMaternal = &personRequest.LastNameMaternal
	params.BirthDate = &personRequest.BirthDate
	params.Gender = &personRequest.Gender
	params.BusinessName = &personRequest.BusinessName
	params.TradeName = &personRequest.TradeName
	params.ConstitutionDate = &personRequest.ConstitutionDate
	params.RepresentativeName = &personRequest.RepresentativeName
	params.RepresentativeDocument = &personRequest.RepresentativeDocument

	return params
}
