package locker_register

type VehicleLockerRegister struct {
	Placa                   string  `json:"placa"`
	Antt                    string  `json:"antt"`
	Rastreador              string  `json:"rastreador"`
	Uf                      string  `json:"uf"`
	CadastroFull            *string `json:"cadastroFull"`
	TipoVeiculo             string  `json:"tipoVeiculo"`
	TipoVinculoProfissional string  `json:"tipoVinculoProfissional"`
}

const (
	CAVALO_MECANICO int = 2
	CARRETA             = 1
)

func NewVehicleLockerRegister() *VehicleLockerRegister {
	return &VehicleLockerRegister{}
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
