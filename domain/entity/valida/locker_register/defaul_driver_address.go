package locker_register

import "strconv"

type DefaultDriverAddress struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Cidade      string `json:"cidade"`
	Numero      string `json:"numero"`
	Uf          string `json:"uf"`
	Referencia  string `json:"referencia"`
}

func NewDefaultDriverAddress() *DefaultDriverAddress {
	dda := &DefaultDriverAddress{
		Cep:         strconv.Itoa(85812030),
		Logradouro:  "Rua Minas Gerais",
		Complemento: "11º Andar. Sala 1105 e 1106",
		Bairro:      "Centro",
		Cidade:      "Cascavel",
		Numero:      "1932",
		Uf:          "PR",
		Referencia:  "",
	}

	return dda
}
