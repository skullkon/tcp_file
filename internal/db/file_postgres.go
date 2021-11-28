package db

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/skullkon/tcp_file/internal/models"
	"log"
	"net"
)

type FilePostgres struct {
	db *sqlx.DB
}

func NewFilePostgres(db *sqlx.DB) *FilePostgres {
	return &FilePostgres{db: db}
}

func (f *FilePostgres) Add(conn net.Conn, uuid string, filePath string) {
	_, err := f.db.Query("INSERT INTO files (uuid,filepath) VALUES (?,?)", uuid, filePath)
	if err != nil {
		return
	}

}

func (f *FilePostgres) ReadAll(conn net.Conn) []models.Entity {
	var files []models.Entity
	var rows, err = f.db.Query("select * from files")
	if err != nil {
		log.Println(err)
		return nil
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)
	for rows.Next() {
		var file models.Entity
		err := rows.Scan(&file.UUID, &file.FilePath)
		if err != nil {
			log.Println(err)
			return nil
		}
		files = append(files, file)
	}
	log.Println(files)
	return files
}
