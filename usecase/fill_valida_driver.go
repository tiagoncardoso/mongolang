package usecase

import (
	"ccovdata/domain/entity"
)

type FillValidaDriver struct {
	Repository entity.CcovRepository
}

func NewFillValidaDriver(repository entity.CcovRepository) *FillValidaDriver {
	return &FillValidaDriver{Repository: repository}
}

func (dr *)
