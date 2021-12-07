package valida

import (
	"ccovdata/domain/entity/valida/locker_register"
	"time"
)

type VehicleRegister struct {
	Id                int32                                `json:"id"`
	TipoVinculo       int16                                `json:"tipo_vinculo_id"`
	TipoVeiculoId     int16                                `json:"tipo_Veiculo_id"`
	ProprietarioId    int32                                `json:"proprietarioId"`
	VeiculoRenovadoId int8                                 `json:"veiculo_renovado_id"`
	Placa             string                               `json:"placa"`
	CadastroGaveta    locker_register.DriverLockerRegister `json:"cadastro_gaveta"`
	Criacao           time.Time                            `json:"criacao"`
	ImgCrlv           string                               `json:"img_crlv"`
	Uf                string                               `json:"uf"`
}
