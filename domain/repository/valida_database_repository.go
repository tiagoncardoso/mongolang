package repository

import (
	"ccovdata/domain/entity/valida"
	"database/sql"
)

type ValidaDatabaseRepository struct {
	db *sql.DB
}

func NewValidaDatabaseRepository(db *sql.DB) *ValidaDatabaseRepository {
	return &ValidaDatabaseRepository{db: db}
}

func (vdr *ValidaDatabaseRepository) InsertDriver(dr *valida.DriverRegister) (int64, error) {
	stmt, err := vdr.db.Prepare(`
			INSERT INTO profissional(tipo_vinculo_id, profissional_renovado_id, nome, cpf, cadastro_gaveta, criacao, img_cnh, img_rg, uf) VALUES
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
