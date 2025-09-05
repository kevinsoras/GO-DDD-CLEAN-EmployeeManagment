// domain/factories/person_params.go
package factories

import (
	"time"

	"github.com/kevinsoras/employee-management/shared/domain/value_objects"
)

// PersonFactoryParams - Parámetros type-safe EN EL DOMINIO
type PersonFactoryParams struct {
	Type    value_objects.PersonType
	Email   value_objects.Email
	Phone   value_objects.Phone
	Address string
	Country string

	DocumentNumber string

	// Campos específicos de Natural
	FirstName        *string
	LastNamePaternal *string
	LastNameMaternal *string
	BirthDate        *time.Time
	Gender           *string

	// Campos específicos de Juridical
	BusinessName           *string
	TradeName              *string
	ConstitutionDate       *time.Time
	RepresentativeName     *string
	RepresentativeDocument *string
}
