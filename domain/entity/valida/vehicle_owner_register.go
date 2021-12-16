package valida

type VehicleOwner struct {
	Id             int     `json:"id"`
	Nome           string  `json:"nome"`
	NomeMae        string  `json:"nomeMae"`
	DataNascimento string  `json:"dataNascimento"`
	Cpf            string  `json:"cpf"`
	Cnpj           string  `json:"cnpj"`
	Rg             string  `json:"rg"`
	UfRg           string  `json:"ufRg"`
	Fone           string  `json:"fone"`
	Complemento    *string `json:"complemento"`
}

func NewVehicleOwner() *VehicleOwner {
	nvo := &VehicleOwner{
		Id:          0,
		Complemento: nil,
	}

	return nvo
}
