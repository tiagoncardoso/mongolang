package entity

type CcovRepository interface {
	GetData(lim int16) error
}
