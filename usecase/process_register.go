package usecase

import (
	"ccovdata/domain/factory"
)

type RegisterProcess struct {
	Repository factory.DriverRegisterExternalRepository
}

func NewRegisterProcess(repository factory.DriverRegisterExternalRepository) *RegisterProcess {
	return &RegisterProcess{Repository: repository}
}
