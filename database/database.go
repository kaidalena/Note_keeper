package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"note_keeper/config"
)

const (
	user_tabel  = "public.users"
	notes_tabel = "public.Notes"
)

var (
	db *sql.DB
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func CreateConn(host, user, user_pwd, dbname string, port int) (*sql.DB, error) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, user_pwd, dbname)

	db, err := sql.Open("postgres", psqlconn)

	return db, err
}

func GetConn() *sql.DB {
	var err error
	if db == nil {
		db, err = CreateConn(config.DB_conf.Host, config.DB_conf.User, config.DB_conf.Password, config.DB_conf.DBname, config.DB_conf.Port)
		CheckError(err)
	}

	return db
}
