package valida

import (
	"time"
)

const (
	ADEQUADO = 1
)

type ResultRegister struct {
	Id                int64     `json:"id"`
	CadastroId        int64     `json:"cadastro_id"`
	CodigoLiberacaoId int       `json:"codigo_liberacao_id"`
	UsuarioId         int64     `json:"usuario_id"`
	TipoResultado     string    `json:"tipo_resultado"`
	Situacao          string    `json:"situacao"`
	InicioLiberacao   string    `json:"inicio_liberacao"`
	FimLiberacao      string    `json:"fim_liberacao"`
	Criacao           time.Time `json:"criacao"`
}

func NewResultRegister(cadastroId int64, code int) *ResultRegister {
	return &ResultRegister{
		CadastroId:        cadastroId,
		CodigoLiberacaoId: code,
		Criacao:           time.Now(),
	}
}

func (res *ResultRegister) SetSituacao(score int) {
	if score <= 10 {
		res.Situacao = "DIVERGENTE"
		res.TipoResultado = "REVERSIVEL"
	} else {
		res.Situacao = "ADEQUADO"
		res.TipoResultado = "TOTAL"
	}
}

func (res *ResultRegister) SetValidade(score int, creation time.Time, validity time.Time) {
	if score > 10 {
		res.InicioLiberacao = creation.Format("YYYY-MM-DD hh:mm")
		res.InicioLiberacao = validity.Format("YYYY-MM-DD hh:mm")
	}
}
