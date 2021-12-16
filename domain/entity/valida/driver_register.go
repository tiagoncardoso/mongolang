package valida

import (
	"ccovdata/domain/entity/valida/locker_register"
	"errors"
	"time"
)

type DriverRegister struct {
	Id             int                                  `json:"id"`
	Name           string                               `json:"nome"`
	Cpf            string                               `json:"cpf"`
	Criacao        time.Time                            `json:"criacao"`
	CadastroGaveta locker_register.DriverLockerRegister `json:"cadastroGaveta"`
	ImgCnh         string                               `json:"img_cng"`
	ImgRg          string                               `json:"img_rg"`
	Uf             string                               `json:"uf"`
}

func NewDriverRegister(name string, cpf string, criacao time.Time, uf string) (*DriverRegister, error) {
	dr := &DriverRegister{
		Id:      0,
		Name:    name,
		Cpf:     cpf,
		Criacao: criacao,
		ImgCnh:  "file_chn_mock.jpg",
		ImgRg:   "file_rg_mock.jpg",
		Uf:      uf,
	}

	err := dr.IsValid()

	if err != nil {
		return nil, err
	}

	return dr, nil
}

func (dr *DriverRegister) IsValid() error {
	if dr.Cpf == "" {
		return errors.New("The driver register needs contains a valid CPF.")
	}

	return nil
}
