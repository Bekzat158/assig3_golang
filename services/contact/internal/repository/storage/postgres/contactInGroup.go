package postgres

import (
	"github.com/google/uuid"

	"assig3/services/contact/internal/domain/contact"
)

func (r *Repository) CreateContactIntoGroup(groupID uuid.UUID, in ...*contact.Contact) ([]*contact.Contact, error) {
	panic("just do something")
}

func (r *Repository) DeleteContactFromGroup(groupID, contactID uuid.UUID) error {
	panic("just do")
}

func (r *Repository) AddContactsToGroup(groupID uuid.UUID, contactIDs ...uuid.UUID) error {
	panic("just do")
}
