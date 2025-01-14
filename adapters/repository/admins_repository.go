package repository

import (
	"database/sql"
	ports "guguzaza-users/ports/repository"
)

type AdminsRepository struct {
	db *sql.DB
}

func NewAdminsRepository(db *sql.DB) ports.AdminsRepositoryPort {
	return nil // todo
}

func (ar *AdminsRepository) CreateAdminUser() {}
