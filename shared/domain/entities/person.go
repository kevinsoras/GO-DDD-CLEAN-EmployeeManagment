package entities

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/kevinsoras/employee-management/shared/domain/value_objects"
)

type Person struct {
	ID        string // UUID
	Type      value_objects.PersonType
	Email     value_objects.Email
	Phone     value_objects.Phone
	Address   string
	Country   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewPerson(pt value_objects.PersonType, email value_objects.Email, phone value_objects.Phone, address, country string) *Person {
	u7, err := uuid.NewV7()
	if err != nil {
		log.Fatalf("Error al generar ID: %v", err)
	}
	return &Person{
		ID:        u7.String(),
		Type:      pt,
		Email:     email,
		Phone:     phone,
		Address:   address,
		Country:   country,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
