package database

import (
	"github.com/jmoiron/sqlx"
)

type DB struct {
	constr string
}

func New(conn string) DB {
	return DB{conn}
}

func (d DB) conn() (*sqlx.DB, error) {
	conn, err := sqlx.Connect("postgres", d.constr)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
