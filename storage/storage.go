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
	once sync.Once
)

func NewPostgresDB() {
	once.Do(func() {

		// Lo que este aqui solo se ejecutara 1 vez
		var err error
		db, err = sql.Open("postgres", "postgres://jhoguer:jhon198615@localhost:5432/godb?sslmode=disable")
		if err != nil {
			log.Fatalf("can't open db: %v", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("can't do ping: %v", err)
		}

		fmt.Println("Conectado a Postgres")
	})
}

// Pool return a unique instace of db
func Pool() *sql.DB {
	return db
}
