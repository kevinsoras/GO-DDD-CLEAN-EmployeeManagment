package entities

import (
	"errors"
	"time"
)

type JuridicalPerson struct {
	PersonID               string // FK a Person
	DocumentNumber         string
	BusinessName           string
	TradeName              string
	ConstitutionDate       time.Time
	RepresentativeName     string
	RepresentativeDocument string
}

// Constructor
func NewJuridicalPerson(personID string, docNum string, businessName *string,
	tradeName *string, repName *string, repDoc *string, constDate *time.Time) (*JuridicalPerson, error) {

	j := &JuridicalPerson{
		PersonID:       personID,
		DocumentNumber: docNum,
	}
	if businessName != nil {
		j.BusinessName = *businessName
	}
	if tradeName != nil {
		j.TradeName = *tradeName
	}
	if repName != nil {
		j.RepresentativeName = *repName
	}
	if repDoc != nil {
		j.RepresentativeDocument = *repDoc
	}
	if constDate != nil {
		j.ConstitutionDate = *constDate
	}

	// Validación
	if err := j.Validate(); err != nil {
		return nil, err
	}

	return j, nil
}

// Validate - valida campos requeridos y reglas de negocio
func (n *JuridicalPerson) Validate() error {
	if n.PersonID == "" {
		return errors.New("person ID is required")
	}
	if n.BusinessName == "" {
		return errors.New("business name is required")
	}
	if n.DocumentNumber == "" {
		return errors.New("document number is required")
	}
	if n.RepresentativeName == "" {
		return errors.New("representative name is required")
	}
	if n.RepresentativeDocument == "" {
		return errors.New("representative document is required")
	}
	if n.TradeName == "" {
		return errors.New("trade name is required")
	}
	if n.ConstitutionDate.IsZero() {
		return errors.New("constitution date is required")
	}
	return nil
}
