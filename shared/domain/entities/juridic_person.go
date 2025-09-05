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
	if len(n.BusinessName) > 100 {
		return errors.New("business name too long")
	}
	if n.DocumentNumber == "" {
		return errors.New("document number is required")
	}
	// Validación de RUC peruano: 11 dígitos numéricos
	if len(n.DocumentNumber) != 11 {
		return errors.New("el RUC debe tener 11 dígitos")
	}
	for _, c := range n.DocumentNumber {
		if c < '0' || c > '9' {
			return errors.New("el RUC solo puede contener números")
		}
	}
	if n.RepresentativeName == "" {
		return errors.New("representative name is required")
	}
	if len(n.RepresentativeName) > 100 {
		return errors.New("representative name too long")
	}
	if n.RepresentativeDocument == "" {
		return errors.New("representative document is required")
	}
	if len(n.RepresentativeDocument) > 20 {
		return errors.New("representative document too long")
	}
	if n.TradeName == "" {
		return errors.New("trade name is required")
	}
	if len(n.TradeName) > 100 {
		return errors.New("trade name too long")
	}
	if n.ConstitutionDate.IsZero() {
		return errors.New("constitution date is required")
	}
	// Fecha de constitución no puede ser futura ni antes de 1900
	if n.ConstitutionDate.After(time.Now()) {
		return errors.New("constitution date cannot be in the future")
	}
	if n.ConstitutionDate.Before(time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC)) {
		return errors.New("constitution date cannot be before 1900")
	}
	return nil
}
