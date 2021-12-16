package valida

import (
	"errors"
	"time"
)

const (
	ADEQUADO = 1
)

type ResultRegister struct {
	Id                int32     `json:"id"`
	CadastroId        int32     `json:"cadastro_id"`
	CodigoLiberacaoId int32     `json:"codigo_liberacao_id"`
	UsuarioId         int32     `json:"usuario_id"`
	TipoResultado     string    `json:"tipo_resultado"`
	Situacao          string    `json:"situacao"`
	InicioLiberacao   time.Time `json:"inicio_liberacao"`
	FimLiberacao      time.Time `json:"fim_liberacao"`
	Criacao           time.Time `json:"criacao"`
}

func (result ResultRegister) NewResultRegister(cadastro MainRegister, tipoResultado string, situacao string, criacao time.Time) (error, ResultRegister) {
	result.CadastroId = cadastro.Id
	result.TipoResultado = tipoResultado
	result.Situacao = situacao
	result.Criacao = criacao

	err := IsValidResult(result)

	if err != nil {
		return err, ResultRegister{}
	}

	return nil, result
}

func IsValidResult(result ResultRegister) error {
	if result.CadastroId == 0 {
		return errors.New("Não foi informado profissional ou veículo")
	}

	return nil
}
