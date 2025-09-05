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
	if n.FirstName == "" {
		return errors.New("firstName es obligatorio")
	}
	if n.LastNamePaternal == "" {
		return errors.New("lastNamePaternal es obligatorio")
	}
	if n.Gender != "M" && n.Gender != "F" && n.Gender != "O" {
		return errors.New("gender debe ser M, F o O")
	}
	if n.BirthDate.IsZero() {
		return errors.New("birthDate es obligatorio")
	}
	return nil
}
