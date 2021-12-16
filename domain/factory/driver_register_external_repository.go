package factory

import (
	"ccovdata/domain/entity/ccov"
	"ccovdata/domain/repository"
)

type DriverRegisterExternalRepository interface {
	ProcessCcovRegisterExternal(repo *repository.CcovRegisterRepository) ([]*ccov.DriverRegisterExternal, error)
}
