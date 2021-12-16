package valida

import (
	"ccovdata/domain/entity/valida/locker_register"
	"errors"
	"time"
)

type VehicleRegister struct {
	Id                int                                  `json:"id"`
	TipoVinculo       int                                  `json:"tipo_vinculo_id"`
	TipoVeiculoId     int                                  `json:"tipo_Veiculo_id"`
	ProprietarioId    int                                  `json:"proprietarioId"`
	VeiculoRenovadoId int                                  `json:"veiculo_renovado_id"`
	Placa             string                               `json:"placa"`
	CadastroGaveta    locker_register.DriverLockerRegister `json:"cadastro_gaveta"`
	Criacao           time.Time                            `json:"criacao"`
	ImgCrlv           string                               `json:"img_crlv"`
	Uf                string                               `json:"uf"`
}

const (
	AUTONOMO int = 1
	AGREGADO     = 2
	FROTA        = 3
)

func NewVehicleRegister(placa string, uf string) (*VehicleRegister, error) {
	ve := &VehicleRegister{
		Placa: placa,
		Uf:    uf,
	}

	err := ve.IsValid()

	if err != nil {
		return nil, err
	}

	return ve, nil
}

func (ve *VehicleRegister) SetTipoVinculo(tipoVinculo string) {
	switch tipoVinculo {
	case "Terceiro":
		ve.TipoVinculo = AUTONOMO
	case "Agregado":
		ve.TipoVinculo = AGREGADO
	case "Frota":
		ve.TipoVinculo = FROTA
	}
}

func (ve *VehicleRegister) IsValid() error {
	if ve.Placa == "" {
		return errors.New("Error on vehicle register. Plates don't be null.")
	}

	return nil
}
