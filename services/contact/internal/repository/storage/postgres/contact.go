package postgres

import (
	"github.com/google/uuid"

	"assig3/pkg/type/queryParameter"
	"assig3/services/contact/internal/domain/contact"
)

func (r *Repository) CreateContact(contacts ...*contact.Contact) ([]*contact.Contact, error) {
	panic("just do something")
}

func (r *Repository) UpdateContact(ID uuid.UUID, updateFn func(c *contact.Contact) (*contact.Contact, error)) (*contact.Contact, error) {
	panic("just do som")
}

func (r *Repository) DeleteContact(ID uuid.UUID) error {
	panic("just do som")
}

func (r *Repository) ListContact(parameter queryParameter.QueryParameter) ([]*contact.Contact, error) {
	panic("just do som")
}

func (r *Repository) ReadContactByID(ID uuid.UUID) (response *contact.Contact, err error) {
	panic("just do som")
}

func (r *Repository) CountContact() (uint64, error) {
	panic("just do som")
}
