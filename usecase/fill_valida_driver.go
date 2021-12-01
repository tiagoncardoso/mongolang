package usecase

import "ccovdata/entity"

type FillValidaDriver struct {
	Repository entity.CcovRepository
}

func NewFillValidaDriver(repository entity.CcovRepository) *FillValidaDriver {
	return &FillValidaDriver{Repository: repository}
}