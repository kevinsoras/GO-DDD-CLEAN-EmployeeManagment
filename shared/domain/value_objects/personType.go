package value_objects

import (
	"fmt"
	"strings"
)

type PersonType string

// Definimos solo las constantes
const (
	Natural   PersonType = "NATURAL"
	Juridical PersonType = "JURIDICAL"
)

// validPersonTypes se genera automáticamente en init()
var validPersonTypes map[PersonType]struct{}

func init() {
	validPersonTypes = make(map[PersonType]struct{})
	// Aquí defines UNA SOLA VEZ los tipos válidos
	validPersonTypes[Natural] = struct{}{}
	validPersonTypes[Juridical] = struct{}{}
}

func NewPersonType(input string) (PersonType, error) {
	if input == "" {
		return "", fmt.Errorf("person type cannot be empty")
	}

	normalizedInput := strings.TrimSpace(strings.ToUpper(input))
	personType := PersonType(normalizedInput)

	if _, isValid := validPersonTypes[personType]; !isValid {
		return "", fmt.Errorf("invalid person type: %s", input)
	}

	return personType, nil
}
