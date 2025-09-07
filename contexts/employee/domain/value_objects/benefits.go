package value_objects

import "errors"

// Benefits es un Value Object que representa los beneficios laborales calculados.
// Es inmutable y se valida en su creación.
type Benefits struct {
	cts           float64
	gratification float64
	vacationDays  int
}

// NewBenefits es el constructor para el Value Object Benefits.
// Asegura que los valores sean válidos antes de crear el objeto.
func NewBenefits(cts, gratification float64, vacationDays int) (Benefits, error) {
	if cts < 0 {
		return Benefits{}, errors.New("el valor de CTS no puede ser negativo")
	}
	if gratification < 0 {
		return Benefits{}, errors.New("el valor de la gratificación no puede ser negativo")
	}
	if vacationDays < 0 {
		return Benefits{}, errors.New("los días de vacaciones no pueden ser negativos")
	}

	return Benefits{
		cts:           cts,
		gratification: gratification,
		vacationDays:  vacationDays,
	}, nil
}

// CTS devuelve el valor del Compensación por Tiempo de Servicios.
func (b Benefits) CTS() float64 {
	return b.cts
}

// Gratification devuelve el valor de la gratificación.
func (b Benefits) Gratification() float64 {
	return b.gratification
}

// VacationDays devuelve los días de vacaciones acumulados.
func (b Benefits) VacationDays() int {
	return b.vacationDays
}

// Equals compara si dos Value Objects Benefits son iguales.
func (b Benefits) Equals(other Benefits) bool {
	return b.cts == other.cts && b.gratification == other.gratification && b.vacationDays == other.vacationDays
}
