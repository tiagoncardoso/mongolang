package valida

import (
	"ccovdata/domain/entity/valida/locker_register"
	"errors"
	"time"
)

type DriverRegister struct {
	Id             int32                                 `json:"id"`
	Name           string                                `json:"nome"`
	Cpf            string                                `json:"cpf"`
	Criacao        time.Time                             `json:"criacao"`
	CadastroGaveta locker_register.VehicleLockerRegister `json:"cadastroGaveta"`
	ProprietarioId VehicleOwner                          `json:"proprietario_id"`
	ImgCnh         string                                `json:"img_cng"`
	ImgRg          string                                `json:"img_rg"`
	ImgEndereco    string                                `json:"img_endereco"`
	Uf             string                                `json:"uf"`
}

func (dr *DriverRegister) IsValid() error {
	if dr.Cpf == "" {
		return errors.New("The driver register needs contains a valid CPF.")
	}

	return nil
}
