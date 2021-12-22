package locker_register

type VehicleLockerRegister struct {
	Placa                   string  `json:"placa"`
	Antt                    string  `json:"antt"`
	Rastreador              int     `json:"rastreador"`
	Uf                      string  `json:"uf"`
	CadastroFull            *string `json:"cadastroFull"`
	TipoVeiculo             string  `json:"tipoVeiculo"`
	TipoVinculoProfissional string  `json:"tipoVinculoProfissional"`
}

func NewVehicleLockerRegister() *VehicleLockerRegister {
	return &VehicleLockerRegister{
		Rastreador: 0,
	}
}

func (vlr *VehicleLockerRegister) SetVehicleData(placa string, antt string, uf string) {
	vlr.Placa = placa
	vlr.Antt = antt
	vlr.Uf = uf
	vlr.CadastroFull = nil
}

func (vlr *VehicleLockerRegister) SetVehicleType(Vehicletype string) {
	vlr.TipoVeiculo = Vehicletype
}
