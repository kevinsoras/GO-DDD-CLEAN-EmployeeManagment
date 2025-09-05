package value_objects

import (
	"fmt"
	"regexp"
)

type Email string

func NewEmail(e string) (Email, error) {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !re.MatchString(e) {
		return "", fmt.Errorf("invalid email: %s", e)
	}
	return Email(e), nil
}
