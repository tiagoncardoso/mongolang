package valida

import "time"

type DemandRegister struct {
	Id               int       `json:"id"`
	CadastroId       int64     `json:"cadastro_id"`
	UsuarioId        int       `json:"usuario_id"`
	EmpresaId        int       `json:"empresa_id"`
	Criacao          time.Time `json:"criacao"`
	PrimeiroCadastro int       `json:"primeiro_cadastro"`
	DeFormulario     int       `json:"de_formulario"`
}

func NewDemandRegister(cadastroId int64, empresaId int, createdAt time.Time) *DemandRegister {
	return &DemandRegister{
		CadastroId:       cadastroId,
		UsuarioId:        2343,
		EmpresaId:        empresaId,
		Criacao:          createdAt,
		PrimeiroCadastro: 1,
		DeFormulario:     0,
	}
}
