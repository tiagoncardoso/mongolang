package valida

import (
	"database/sql"
	"time"
)

const (
	ADEQUADO = 1
)

type ResultRegister struct {
	Id                int64         `json:"id"`
	CadastroId        int64         `json:"cadastro_id"`
	CodigoLiberacaoId sql.NullInt64 `json:"codigo_liberacao_id"`
	UsuarioId         int64         `json:"usuario_id"`
	TipoResultado     string        `json:"tipo_resultado"`
	Situacao          string        `json:"situacao"`
	InicioLiberacao   string        `json:"inicio_liberacao"`
	FimLiberacao      string        `json:"fim_liberacao"`
	Criacao           time.Time     `json:"criacao"`
}

func NewResultRegister(cadastroId int64, createdAt time.Time) *ResultRegister {
	return &ResultRegister{
		CadastroId:        cadastroId,
		Criacao:           createdAt,
		CodigoLiberacaoId: sql.NullInt64{},
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
		res.InicioLiberacao = creation.Format("2006-01-02 15:04:05")
		res.FimLiberacao = validity.Format("2006-01-02 15:04:05")
	}
}

func (res *ResultRegister) SetCode(code int64) {
	res.CodigoLiberacaoId = sql.NullInt64{
		Int64: code,
		Valid: true,
	}
}
