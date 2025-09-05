package value_objects

import "fmt"

type Phone string

func NewPhone(p string) (Phone, error) {
	if len(p) < 6 || len(p) > 20 {
		return "", fmt.Errorf("invalid phone number: %s", p)
	}
	return Phone(p), nil
}
