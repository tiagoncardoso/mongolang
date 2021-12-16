package valida

import (
	"time"
)

type MainRegister struct {
	Id               int32           `json:"id"`
	Profissional     DriverRegister  `json:"profissionalId"`
	Motor            VehicleRegister `json:"motor"`
	Carreta1         VehicleRegister `json:"carreta1"`
	Carreta2         VehicleRegister `json:"ccarreta2"`
	Carreta3         VehicleRegister `json:"carreta3"`
	ViagemId         TravelRegister  `json:"viagem_id"`
	TipoCadastro     string          `json:"tipo_cadastro"`
	Validado         bool            `json:"validado"`
	ValidadeCadastro time.Time       `json:"validade_cadastro"`
	Status           string          `json:"status"`
	Criacao          time.Time       `json:"criacao"`
}

func NewMainRegister(criacao time.Time) *MainRegister {
	reg := &MainRegister{
		Id:           0,
		TipoCadastro: "PROFISSIONAL_VEICULO",
		Validado:     true,
		Criacao:      criacao,
	}

	return reg
}

func (reg *MainRegister) SetProfissional(pr DriverRegister) {
	reg.Profissional = pr
}
