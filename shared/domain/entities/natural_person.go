package entities

import (
	"errors"
	"time"
)

type NaturalPerson struct {
	PersonID         string
	DocumentNumber   string
	FirstName        string
	LastNamePaternal string
	LastNameMaternal string
	BirthDate        time.Time
	Gender           string // M, F, O
}

// Constructor con validación interna
func NewNaturalPerson(
	personID string,
	documentNumber string,
	firstName *string,
	lastPat *string,
	lastMat *string,
	gender *string,
	birthDate *time.Time,
) (*NaturalPerson, error) {

	n := &NaturalPerson{
		PersonID:       personID,
		DocumentNumber: documentNumber,
	}

	if firstName != nil {
		n.FirstName = *firstName
	}
	if lastPat != nil {
		n.LastNamePaternal = *lastPat
	}
	if lastMat != nil {
		n.LastNameMaternal = *lastMat
	}
	if gender != nil {
		n.Gender = *gender
	}
	if birthDate != nil {
		n.BirthDate = *birthDate
	}

	// Validación
	if err := n.Validate(); err != nil {
		return nil, err
	}

	return n, nil
}

// Validate - valida campos requeridos y reglas de negocio
func (n *NaturalPerson) Validate() error {
	if n.PersonID == "" {
		return errors.New("personID es obligatorio")
	}
	if n.DocumentNumber == "" {
		return errors.New("documentNumber es obligatorio")
	}
	// Validación de DNI peruano: 8 dígitos numéricos
	if len(n.DocumentNumber) != 8 {
		return errors.New("el DNI debe tener 8 dígitos")
	}
	for _, c := range n.DocumentNumber {
		if c < '0' || c > '9' {
			return errors.New("el DNI solo puede contener números")
		}
	}
	if n.FirstName == "" {
		return errors.New("firstName es obligatorio")
	}
	if len(n.FirstName) > 50 {
		return errors.New("firstName demasiado largo")
	}
	if n.LastNamePaternal == "" {
		return errors.New("lastNamePaternal es obligatorio")
	}
	if len(n.LastNamePaternal) > 50 {
		return errors.New("lastNamePaternal demasiado largo")
	}
	if n.LastNameMaternal != "" && len(n.LastNameMaternal) > 50 {
		return errors.New("lastNameMaternal demasiado largo")
	}
	if n.Gender != "M" && n.Gender != "F" && n.Gender != "O" {
		return errors.New("gender debe ser M, F o O")
	}
	if n.BirthDate.IsZero() {
		return errors.New("birthDate es obligatorio")
	}
	// Fecha de nacimiento no puede ser futura ni ridículamente antigua
	if n.BirthDate.After(time.Now()) {
		return errors.New("birthDate no puede ser en el futuro")
	}
	if n.BirthDate.Before(time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC)) {
		return errors.New("birthDate no puede ser antes de 1900")
	}
	return nil
}
