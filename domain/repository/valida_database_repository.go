package repository

import (
	"ccovdata/domain/entity/valida"
	"database/sql"
	"encoding/json"
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
			INSERT INTO profissional(tipo_vinculo_id, profissional_renovado_id, nome, cpf, cadastro_gaveta, criacao, img_cnh, uf) VALUES
			(?, ?, ?, ?, ?, ?, ?, ?)
	`)

	if err != nil {
		return 0, err
	}

	gaveta, _ := json.Marshal(dr.CadastroGaveta)
	createAt := dr.Criacao.Format("2006-01-02 15:04:05")

	var res sql.Result
	res, err = stmt.Exec(
		dr.DriverType,
		dr.Renewed(),
		dr.Name,
		dr.Cpf,
		gaveta,
		createAt,
		dr.DefaultImage(),
		dr.Uf,
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
			INSERT INTO veiculo(tipo_vinculo_id, tipo_veiculo_id, proprietario_id, veiculo_renovado_id, placa, cadastro_gaveta, img_crlv, uf, criacao) VALUES
			(?, ?, ?, ?, ?, ?, ?, ?, ?)
	`)

	if err != nil {
		return 0, err
	}

	gaveta, _ := json.Marshal(veh.CadastroGaveta)

	var res sql.Result
	res, err = stmt.Exec(
		veh.TipoVinculo,
		veh.TipoVeiculoId,
		veh.ProprietarioId,
		veh.Renewed(),
		veh.Placa,
		gaveta,
		veh.DefaultImage(),
		veh.Uf,
		veh.Criacao.Format("2006-01-02 15:04:05"),
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

	carga, _ := json.Marshal(tr.TipoCarga)

	var res sql.Result
	res, err = stmt.Exec(
		tr.OrigemPaisId,
		tr.OrigemUfId,
		tr.OrigemCidadeId,
		tr.DestinoPaisId,
		tr.DestinoUfId,
		tr.DestinoCidadeId,
		carga,
		tr.ValorCarga,
		tr.Criacao.Format("2006-01-02 15:04:05"),
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
			INSERT INTO cadastro(profissional_id, motor_id, carreta1_id, carreta2_id, carreta3_id, viagem_id, tipo_cadastro, validado, validade_cadastro, status, criacao, criminal, plus, protocolo_ccov) 
			VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`)

	if err != nil {
		return 0, err
	}

	vm := getVehicleId(reg.MotorId)
	ca1 := getVehicleId(reg.Carreta1Id)
	ca2 := getVehicleId(reg.Carreta2Id)
	ca3 := getVehicleId(reg.Carreta3Id)

	var res sql.Result
	res, err = stmt.Exec(
		reg.ProfissionalId,
		vm,
		ca1,
		ca2,
		ca3,
		reg.ViagemId,
		reg.TipoCadastro,
		reg.Validado,
		reg.ValidadeCadastro,
		reg.Status,
		reg.Criacao,
		reg.Criminal,
		reg.Plus,
		reg.CcovProtocol,
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
			INSERT INTO resultado(cadastro_id, codigo_liberacao_id, usuario_id, tipo_resultado, situacao, inicio_liberacao, fim_liberacao, criacao) 
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

func (vdr *ValidaDatabaseRepository) GenerateCode() (int64, error) {
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

	var res sql.Result
	res, err = stmt.Exec(
		codeType,
		newCode,
		time.Now(),
	)

	if err != nil {
		return -1, err
	}

	lid, _ := res.LastInsertId()

	return lid, nil
}

func getVehicleId(vehicleId int64) sql.NullInt64 {
	vid := sql.NullInt64{}

	if vehicleId > 0 {
		vid = sql.NullInt64{
			Int64: vehicleId,
			Valid: true,
		}
	}

	return vid
}
