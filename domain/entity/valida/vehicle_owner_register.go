package valida

import "time"

type VehicleOwner struct {
	UfRgId         int32     `json:"uf_rg_id"`
	Cpf            int64     `json:"cpf"`
	Cnpj           int64     `json:"cnpj"`
	Rg             string    `json:"rg"`
	Fone           int64     `json:"fone"`
	Nome           string    `json:"nome"`
	Criacao        time.Time `json:"criacao"`
	NomeMae        string    `json:"nome_mae"`
	DataNascimento time.Time `json:"data_nascimento"`
}
