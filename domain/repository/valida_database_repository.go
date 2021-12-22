package repository

import (
	"ccovdata/domain/entity/valida"
	"database/sql"
	"time"
)

type ValidaDatabaseRepository struct {
	db *sql.DB
}

func NewValidaDatabaseRepository(db *sql.DB) *ValidaDatabaseRepository {
	return &ValidaDatabaseRepository{db: db}
}

func (vdr *ValidaDatabaseRepository) InsertDriver(dr *valida.DriverRegister) (int64, error) {
	stmt, err := vdr.db.Prepare(`
			INSERT INTO profissional(tipo_vinculo_id, profissional_renovado_id, nome, cpf, cadastro_gaveta, criacao, img_cnh, uf, criacao) VALUES
			(?, ?, ?, ?, ?, ?, ?, ?, ?)
	`)

	if err != nil {
		return 0, err
	}

	var res sql.Result
	res, err = stmt.Exec(
		dr.DriverType,
		dr.Renewed(),
		dr.Name,
		dr.Cpf,
		dr.CadastroGaveta,
		dr.DefaultImage(),
		dr.Uf,
		dr.Criacao.Format("YYYY-MM-DD"),
	)

	if err != nil {
		return 0, err
	}

	var lid int64
	lid, err = res.LastInsertId()

	if err != nil {
		return 0, err
	}

	return lid, nil
}

func (vdr *ValidaDatabaseRepository) InsertVehicle(veh *valida.VehicleRegister) (int64, error) {
	stmt, err := vdr.db.Prepare(`
			INSERT INTO veiculo(tipo_vinculo_id, tipo_veiculo_id, proprietario_id, veiculo_renovado_id, placa, cadastro_gaveta, criacao, img_crlv, uf) VALUES
			(?, ?, ?, ?, ?, ?, ?, ?, ?)
	`)

	if err != nil {
		return 0, err
	}

	var res sql.Result
	res, err = stmt.Exec(
		veh.TipoVinculo,
		veh.TipoVeiculoId,
		veh.ProprietarioId,
		veh.Renewed(),
		veh.Placa,
		veh.CadastroGaveta,
		veh.DefaultImage(),
		veh.Uf,
		veh.Criacao.Format("YYYY-MM-DD"),
	)

	if err != nil {
		return 0, err
	}

	var lid int64
	lid, err = res.LastInsertId()

	if err != nil {
		return 0, err
	}

	return lid, nil
}

func (vdr *ValidaDatabaseRepository) InsertTravel(tr *valida.TravelRegister) (int64, error) {
	stmt, err := vdr.db.Prepare(`
			INSERT INTO viagem(origem_pais_id, origem_uf_id, origem_cidade_id, destino_pais_id, destino_uf_id, destino_cidade_id, tipo_carga, valor_carga, criacao) 
			VALUES
			(?, ?, ?, ?, ?, ?, ?, ?, ?)
	`)

	if err != nil {
		return 0, err
	}

	var res sql.Result
	res, err = stmt.Exec(
		tr.OrigemPaisId,
		tr.OrigemUfId,
		tr.OrigemCidadeId,
		tr.DestinoPaisId,
		tr.DestinoUfId,
		tr.DestinoCidadeId,
		tr.TipoCarga,
		tr.ValorCarga,
		tr.Criacao,
	)

	if err != nil {
		return 0, err
	}

	var lid int64
	lid, err = res.LastInsertId()

	if err != nil {
		return 0, err
	}

	return lid, nil
}

func (vdr *ValidaDatabaseRepository) InsertRegister(reg *valida.Register) (int64, error) {
	stmt, err := vdr.db.Prepare(`
			INSERT INTO cadastro(profissional_id, motor_id, carreta1_id, carreta2_id, carreta3_id, viagem_id, tipo_cadastro, validado, validade_cadastro, status, criacao, criminal, plus) 
			VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`)

	if err != nil {
		return 0, err
	}

	var res sql.Result
	res, err = stmt.Exec(
		reg.ProfissionalId,
		reg.MotorId,
		reg.Carreta1Id,
		reg.Carreta2Id,
		reg.Carreta3Id,
		reg.ViagemId,
		reg.TipoCadastro,
		reg.Validado,
		reg.ValidadeCadastro,
		reg.Status,
		reg.Criacao,
		reg.Criminal,
		reg.Plus,
	)

	if err != nil {
		return 0, err
	}

	var lid int64
	lid, err = res.LastInsertId()

	if err != nil {
		return 0, err
	}

	return lid, nil
}

func (vdr *ValidaDatabaseRepository) InsertResultRegister(result *valida.ResultRegister) (int64, error) {
	stmt, err := vdr.db.Prepare(`
			INSERT INTO resultado(cadastro, codigo_liberacao_id, usuario_id, tipo_resultado, situacao, inicio_liberacao, fim_liberacao, criacao) 
			VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`)

	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(
		result.CadastroId,
		result.CodigoLiberacaoId,
		1,
		result.TipoResultado,
		result.Situacao,
		result.InicioLiberacao,
		result.FimLiberacao,
		result.Criacao,
	)

	if err != nil {
		return 0, err
	}

	var lid int64
	lid, err = res.LastInsertId()

	if err != nil {
		return 0, err
	}

	return lid, nil
}

func (vdr *ValidaDatabaseRepository) GenerateCode() (int, error) {
	var newCode int
	codeType := "L"

	err := vdr.db.QueryRow(`SELECT (counter + 1) AS code FROM codes WHERE type = ? ORDER BY id DESC LIMIT 1`, codeType).Scan(&newCode)
	if err != nil {
		return -1, err
	}

	stmt, _ := vdr.db.Prepare(`
			INSERT INTO codes(type, counter, criacao) VALUES
			(?, ?, ?)
	`)

	_, err = stmt.Exec(
		codeType,
		newCode,
		time.Now(),
	)

	if err != nil {
		return -1, err
	}

	return newCode, nil
}
