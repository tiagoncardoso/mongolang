package repository

import (
	"database/sql"
)

type PortalDatabaseRepository struct {
	db *sql.DB
}

func NewPortalDatabaseRepository(db *sql.DB) *PortalDatabaseRepository {
	return &PortalDatabaseRepository{db: db}
}

func (pdr *PortalDatabaseRepository) FindCompanyId(cnpj string) (int, error) {
	var id int

	err := pdr.db.QueryRow(`SELECT id FROM empresa WHERE cnpj = ?`, cnpj).Scan(&id)

	if err != nil {
		return -1, err
	}

	return id, nil
}
