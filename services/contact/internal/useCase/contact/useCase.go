package contact

import (
	"assig3/services/contact/internal/useCase/adapters/storage"
)

type UseCase struct {
	adapterStorage storage.Contact
	options        Options
}

type Options struct{}

func New(storage storage.Contact, options Options) *UseCase {
	var uc = &UseCase{
		adapterStorage: storage,
	}
	uc.SetOptions(options)
	return uc
}

func (uc *UseCase) SetOptions(options Options) {
	if uc.options != options {
		uc.options = options
	}
}
