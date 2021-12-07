package valida

import (
	"errors"
	"time"
)

type MainRegister struct {
	Profissional     int32     `json:"profissionalId"`
	Motor            int32     `json:"motor"`
	Carreta1         int32     `json:"carreta1"`
	Carreta2         int32     `json:"ccarreta2"`
	Carreta3         int32     `json:"carreta3"`
	ViagemId         int32     `json:"viagem_id"`
	TipoCadastro     string    `json:"tipo_cadastro"`
	Validado         bool      `json:"validado"`
	ValidadeCadastro time.Time `json:"validade_cadastro"`
	Status           string    `json:"status"`
	Criacao          time.Time `json:"criacao"`
}

func (reg MainRegister) NewMainRegister(profissional DriverRegister, motor VehicleRegister, carreta1 VehicleRegister, carreta2 VehicleRegister, carreta3 VehicleRegister, viagem TravelRegister, tipoCadastro string, validado bool, validadeCadastro time.Time, status string, criacao time.Time) (MainRegister, error) {
	reg.Profissional = profissional.Id
	reg.Motor = motor.Id
	reg.Carreta1 = carreta1.Id
	reg.Carreta2 = carreta2.Id
	reg.Carreta3 = carreta3.Id
	reg.ViagemId = viagem.Id
	reg.TipoCadastro = tipoCadastro
	reg.Validado = validado
	reg.ValidadeCadastro = validadeCadastro
	reg.Status = status
	reg.Criacao = criacao

	err := IsValid(reg)

	if err != nil {
		return MainRegister{}, err
	}

	return reg, nil
}

func IsValid(reg MainRegister) error {
	if reg.Profissional == 0 && reg.Motor == 0 && reg.Carreta1 == 0 && reg.Carreta2 == 0 && reg.Carreta3 == 0 {
		return errors.New("Não foi informado profissional ou veículo")
	}

	return nil
}
