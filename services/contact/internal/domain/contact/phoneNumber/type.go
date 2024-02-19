package phoneNumber

import (
	"strings"
	"errors"
)
var (
	ErrPhoneNumberRequired = errors.New("phone number is required")
)

type PhoneNumber struct {
	value string
}

func (p PhoneNumber) String() string {
	return p.value
}

func New(phone string) (*PhoneNumber, error) {
    if len(phone) == 0 {
        return nil, ErrPhoneNumberRequired
    }

    return &PhoneNumber{
        value: getNumbers(phone),
    }, nil
}

func (p PhoneNumber) Equal(phoneNumber PhoneNumber) bool {
	return p.value == phoneNumber.value
}

func (p PhoneNumber) IsEmpty() bool {
	return len(strings.TrimSpace(p.value)) == 0
}
