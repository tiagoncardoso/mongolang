package locker_register

import "strconv"

type DriverLockerRegister struct {
	Cpf                 string                `json:"cpf"`
	Nome                string                `json:"nome"`
	Rg                  string                `json:"rg"`
	UfRg                string                `json:"ufRg"`
	Fones               []int                 `json:"fones"`
	Email               *string               `json:"email"`
	Cnh                 string                `json:"cnh"`
	UfCnh               string                `json:"ufCnh"`
	NomeMae             string                `json:"nomeMae"`
	DataNascimento      string                `json:"dataNascimento"`
	ValidadeCnh         string                `json:"validadeCnh"`
	CategoriaCnh        string                `json:"categoriaCnh"`
	CadastroFull        *string               `json:"cadastroFull"`
	Endereco            *DefaultDriverAddress `json:"endereco"`
	ReferenciaPessoal   []string              `json:"referenciaPessoal"`
	ReferenciaComercial []string              `json:"referenciaComercial"`
	TipoProfissional    string                `json:"tipoProfissional"`
}

func NewLockerRegister() *DriverLockerRegister {
	return &DriverLockerRegister{
		CadastroFull: nil,
		Email:        nil,
	}
}

func (dlr *DriverLockerRegister) SetDriverCategory(category string) {
	dlr.TipoProfissional = category
}

func (dlr *DriverLockerRegister) SetPersonalData(nome string, nomeMae string, dataNascimento string) {
	dlr.Nome = nome
	dlr.NomeMae = nomeMae
	dlr.DataNascimento = dataNascimento
}

func (dlr *DriverLockerRegister) SetDocumentsData(cpf string, cnh string, ufCnh string, rg string, ufRg string, validadeCnh string, categoriaCnh string) {
	dlr.Cpf = cpf
	dlr.Cnh = cnh
	dlr.UfCnh = ufCnh
	dlr.Rg = rg
	dlr.UfRg = ufRg
	dlr.ValidadeCnh = validadeCnh
	dlr.CategoriaCnh = categoriaCnh
}

func (dlr *DriverLockerRegister) SetContactData(fone string) {
	var emptyArr []string
	intPhone, _ := strconv.Atoi(fone)

	dlr.Fones = []int{intPhone}
	dlr.Endereco = NewDefaultDriverAddress()
	dlr.ReferenciaPessoal = emptyArr
	dlr.ReferenciaComercial = emptyArr
}
