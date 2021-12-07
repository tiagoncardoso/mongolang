package locker_register

import "ccovdata/domain/entity/valida"

type VehicleLockerRegister struct {
	Placa            string              `json:"placa"`
	Antt             int32               `json:"antt"`
	Uf               int16               `json:"uf"`
	PessoaFisica     bool                `json:"pessoaFisica"`
	Proprietario     valida.VehicleOwner `json:"proprietario"`
	TipoProprietario string              `json:"tipoProprietario"`
	NumCarretas      int32               `json:"numCarretas"`
}
