package storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var (
	db   *sql.DB
	once sync.Once //Crear el singleton
)

func NewPostgresDB() {
	once.Do(func() {
		var err error
		db, err = sql.Open("postgres", "postgres://saalcazar:5254@localhost:5432/godb?sslmode=disable")
		if err != nil {
			log.Fatalf("can`t open db: %v", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("can`t do ping: %v", err)
		}

		fmt.Println("Conectado a postgres")
	})
}

// Retorna una unica instancia de DB
func Pool() *sql.DB {
	return db
}

func stringToNull(s string) sql.NullString {
	null := sql.NullString{String: s}
	if null.String != "" {
		null.Valid = true
	}
	return null
}
