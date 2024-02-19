package postgres

import (
	"github.com/google/uuid"

	"assig3/pkg/type/queryParameter"
	"assig3/services/contact/internal/domain/group"
)

func (r *Repository) CreateGroup(group *group.Group) (*group.Group, error) {
	panic("just do somethi do something")
}

func (r *Repository) UpdateGroup(ID uuid.UUID, updateFn func(group *group.Group) (*group.Group, error)) (*group.Group, error) {
	panic("just do somethi")
}

func (r *Repository) DeleteGroup(ID uuid.UUID) error {
	panic("just do something")
}

func (r *Repository) ListGroup(parameter queryParameter.QueryParameter) ([]*group.Group, error) {
	panic("just do somethi")
}

func (r *Repository) ReadGroupByID(ID uuid.UUID) (*group.Group, error) {
	panic("just do somethi")
}

func (r *Repository) CountGroup() (uint64, error) {
	panic("just do somethi")
}
