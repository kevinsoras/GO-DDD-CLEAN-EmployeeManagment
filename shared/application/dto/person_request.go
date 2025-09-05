// interfaces/http/dto/person_request.go
package dto

import (
	"time"
)

type PersonRequest struct {
	// Campos base (obligatorios para todos)
	Type           string `json:"type" validate:"required,oneof=NATURAL JURIDICAL"`
	Email          string `json:"email" validate:"required,email"`
	Phone          string `json:"phone" validate:"required"`
	Address        string `json:"address" validate:"required"`
	Country        string `json:"country" validate:"required"`
	DocumentNumber string `json:"documentNumber" validate:"required"`

	// Campos específicos para Natural
	FirstName        string    `json:"firstName" validate:"required_if=Type NATURAL"`
	LastNamePaternal string    `json:"lastNamePaternal" validate:"required_if=Type NATURAL"`
	LastNameMaternal string    `json:"lastNameMaternal" validate:"required_if=Type NATURAL"`
	BirthDate        time.Time `json:"birthDate" validate:"required_if=Type NATURAL"`
	Gender           string    `json:"gender" validate:"required_if=Type NATURAL,omitempty,oneof=M F O"`
	// Campos específicos para Juridical
	BusinessName           string    `json:"businessName" validate:"required_if=Type JURIDICAL"`
	TradeName              string    `json:"tradeName" validate:"required_if=Type JURIDICAL"`
	ConstitutionDate       time.Time `json:"constitutionDate" validate:"required_if=Type JURIDICAL"`
	RepresentativeName     string    `json:"representativeName" validate:"required_if=Type JURIDICAL"`
	RepresentativeDocument string    `json:"representativeDocument" validate:"required_if=Type JURIDICAL"`
}
