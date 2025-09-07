package dto

import (
	"time"

	"github.com/kevinsoras/employee-management/shared/domain/aggregates"
	"github.com/kevinsoras/employee-management/shared/domain/value_objects"
)

type PersonResponse struct {
	ID             string    `json:"id"`
	Type           string    `json:"type"`
	Email          string    `json:"email"`
	Phone          string    `json:"phone"`
	Address        string    `json:"address"`
	Country        string    `json:"country"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	DocumentNumber string    `json:"documentNumber,omitempty"`
	// NATURAL
	FirstName        string    `json:"firstName,omitempty"`
	LastNamePaternal string    `json:"lastNamePaternal,omitempty"`
	LastNameMaternal string    `json:"lastNameMaternal,omitempty"`
	BirthDate        time.Time `json:"birthDate,omitempty"`
	Gender           string    `json:"gender,omitempty"`
	// JURIDICAL
	BusinessName           string    `json:"businessName,omitempty"`
	TradeName              string    `json:"tradeName,omitempty"`
	ConstitutionDate       time.Time `json:"constitutionDate,omitempty"`
	RepresentativeName     string    `json:"representativeName,omitempty"`
	RepresentativeDocument string    `json:"representativeDocument,omitempty"`
}

func NewPersonResponse(agg *aggregates.PersonAggregate) PersonResponse {
	pr := PersonResponse{
		ID:        agg.Person.ID,
		Type:      string(agg.Person.Type),
		Email:     string(agg.Person.Email),
		Phone:     string(agg.Person.Phone),
		Address:   agg.Person.Address,
		Country:   agg.Person.Country,
		CreatedAt: agg.Person.CreatedAt,
		UpdatedAt: agg.Person.UpdatedAt,
	}
	// Usar un factory para poblar los campos según el tipo
	switch agg.Person.Type {
	case value_objects.Natural:
		fillNaturalPersonFields(&pr, agg)
	case value_objects.Juridical:
		fillJuridicalPersonFields(&pr, agg)
	}
	return pr
}

func fillNaturalPersonFields(pr *PersonResponse, agg *aggregates.PersonAggregate) {
	if agg.NaturalPerson == nil {
		return
	}
	pr.FirstName = agg.NaturalPerson.FirstName
	pr.LastNamePaternal = agg.NaturalPerson.LastNamePaternal
	pr.LastNameMaternal = agg.NaturalPerson.LastNameMaternal
	pr.BirthDate = agg.NaturalPerson.BirthDate
	pr.Gender = agg.NaturalPerson.Gender
	pr.DocumentNumber = agg.NaturalPerson.DocumentNumber
}

func fillJuridicalPersonFields(pr *PersonResponse, agg *aggregates.PersonAggregate) {
	if agg.JuridicalPerson == nil {
		return
	}
	pr.BusinessName = agg.JuridicalPerson.BusinessName
	pr.TradeName = agg.JuridicalPerson.TradeName
	pr.ConstitutionDate = agg.JuridicalPerson.ConstitutionDate
	pr.RepresentativeName = agg.JuridicalPerson.RepresentativeName
	pr.RepresentativeDocument = agg.JuridicalPerson.RepresentativeDocument
	pr.DocumentNumber = agg.JuridicalPerson.DocumentNumber
}
