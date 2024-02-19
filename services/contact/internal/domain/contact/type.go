package contact

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"assig3/services/contact/internal/domain/contact/email"
	"assig3/services/contact/internal/domain/contact/phoneNumber"
	"assig3/services/contact/internal/domain/contact/gender"
	"assig3/services/contact/internal/domain/contact/age"
	"assig3/services/contact/internal/domain/contact/name"
	"assig3/services/contact/internal/domain/contact/patronymic"
	"assig3/services/contact/internal/domain/contact/surname"
)


type Contact struct {
	id         uuid.UUID
	createdAt  time.Time
	modifiedAt time.Time

	phoneNumber phoneNumber.PhoneNumber
	email       email.Email

	name       name.Name
	surname    surname.Surname
	patronymic patronymic.Patronymic

	age age.Age

	gender gender.Gender
}

func NewWithID(
	id uuid.UUID,
	createdAt time.Time,
	modifiedAt time.Time,
	phoneNumber phoneNumber.PhoneNumber,
	email email.Email,
	name name.Name,
	surname surname.Surname,
	patronymic patronymic.Patronymic,
	age age.Age,
	gender gender.Gender,
) (*Contact, error) {

	
	if id == uuid.Nil {
		id = uuid.New()
	}

	return &Contact{
		id:          id,
		createdAt:   createdAt.UTC(),
		modifiedAt:  modifiedAt.UTC(),
		phoneNumber: phoneNumber,
		email:       email,
		name:        name,
		surname:     surname,
		patronymic:  patronymic,
		age:         age,
		gender:      gender,
	}, nil
}

func New(
	phoneNumber phoneNumber.PhoneNumber,
	email email.Email,
	name name.Name,
	surname surname.Surname,
	patronymic patronymic.Patronymic,
	age age.Age,
	gender gender.Gender,
) (*Contact, error) {

	var timeNow = time.Now().UTC()
	return &Contact{
		id:          uuid.New(),
		createdAt:   timeNow,
		modifiedAt:  timeNow,
		phoneNumber: phoneNumber,
		email:       email,
		name:        name,
		surname:     surname,
		patronymic:  patronymic,
		age:         age,
		gender:      gender,
	}, nil
}

func (c Contact) ID() uuid.UUID {
	return c.id
}

func (c Contact) CreatedAt() time.Time {
	return c.createdAt
}

func (c Contact) ModifiedAt() time.Time {
	return c.modifiedAt
}

func (c Contact) Email() email.Email {
	return c.email
}

func (c Contact) PhoneNumber() phoneNumber.PhoneNumber {
	return c.phoneNumber
}

func (c Contact) Name() name.Name {
	return c.name
}

func (c Contact) Surname() surname.Surname {
	return c.surname
}

func (c Contact) Patronymic() patronymic.Patronymic {
	return c.patronymic
}

func (c Contact) FullName() string {
	return fmt.Sprintf("%s %s %s", c.surname, c.name, c.patronymic)
}

func (c Contact) Age() age.Age {
	return c.age
}

func (c Contact) Gender() gender.Gender {
	return c.gender
}

func (c Contact) Equal(contact Contact) bool {
	return c.id == contact.id
}
