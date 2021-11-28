package db

import (
	"github.com/jmoiron/sqlx"
	"github.com/skullkon/tcp_file/internal/models"
	"net"
)

type File interface {
	Add(conn net.Conn, uuid string, filePath string)
	ReadAll(conn net.Conn) []models.Entity
}

type Repository struct {
	File
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		File: NewFilePostgres(db),
	}
}
