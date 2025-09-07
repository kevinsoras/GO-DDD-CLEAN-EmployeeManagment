package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/go-playground/validator/v10"
	sharedDomain "github.com/kevinsoras/employee-management/shared/domain"
)

// RequiredIf es una validación personalizada para required_if
func RequiredIf(fl validator.FieldLevel) bool {
	field := fl.Field()
	param := fl.Param() // ejemplo: "Type=NATURAL"
	var key, value string
	n, _ := fmt.Sscanf(param, "Type %s", &value)
	if n == 1 {
		key = "Type"
	} else {
		fmt.Sscanf(param, "Type=%s", &value)
		key = "Type"
	}

	// Buscar el valor de Type en el struct padre, aunque esté anidado
	var typeValue string
	parent := fl.Parent().Interface()
	// Buscar recursivamente el campo Type
	switch v := parent.(type) {
	case interface{ GetType() string }:
		typeValue = v.GetType()
	default:
		rv := reflect.ValueOf(parent)
		if rv.Kind() == reflect.Struct {
			f := rv.FieldByName("Type")
			if f.IsValid() && f.Kind() == reflect.String {
				typeValue = f.String()
			}
		}
	}

	if key == "Type" && typeValue == value {
		if field.Kind() == reflect.String {
			return field.String() != ""
		}
		if field.Type().String() == "time.Time" {
			t, ok := field.Interface().(time.Time)
			return ok && !t.IsZero()
		}
	}
	return true
}

// ValidateAndBind simplifica el parseo y validación de un request
func ValidateAndBind(r *http.Request, dst interface{}) error {
	validate := validator.New()

	// Registrar la validación personalizada required_if si no está registrada
	_ = validate.RegisterValidation("required_if", RequiredIf)

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(dst); err != nil {
		return sharedDomain.NewInvalidInputError(fmt.Sprintf("Cuerpo de la solicitud inválido: %s", err.Error()), err)
	}
	if err := validate.Struct(dst); err != nil {
		return sharedDomain.NewInvalidInputError(fmt.Sprintf("Error de validación: %s", err.Error()), err)
	}
	return nil
}
