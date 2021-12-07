package locker_register

type DriverLockerRegister struct {
	Cpf                 int64   `json:"cpf"`
	Nome                string  `json:"nome"`
	Rg                  string  `json:"rg"`
	UfRg                string  `json:"ufRg"`
	Fones               []int64 `json:"fones"`
	Email               string  `json:"email"`
	Cnh                 int64   `json:"cnh"`
	UfCnh               string  `json:"ufCnh"`
	NomeMae             string  `json:"nomeMae"`
	DataNascimento      string  `json:"dataNascimento"`
	ValidadeCnh         string  `json:"validadeCnh"`
	CategoriaCnh        string  `json:"categoriaCnh"`
	CadastroFull        string  `json:"cadastroFull"`
	Endereco            string  `json:"endereco"`
	ReferenciaPessoal   string  `json:"referenciaPessoal"`
	ReferenciaComercial string  `json:"referenciaComercial"`
	TipoProfissional    string  `json:"tipoProfissional"`
}

func NewLockerRegister() *DriverLockerRegister {
	return &DriverLockerRegister{}
}
