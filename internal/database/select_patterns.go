package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/zen-en-tonal/kanaami/internal/host"
	"github.com/zen-en-tonal/kanaami/internal/patterns"
)

type Selector struct {
	*DB
	contentId []byte
}

func (d *DB) MakeSelector(id []byte) Selector {
	return Selector{d, id}
}

func (s Selector) CreateAllows() (*host.Patterns, error) {
	db, err := s.conn()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	allows, err := selectAllows(db, string(s.contentId))
	if err != nil {
		return nil, err
	}

	return patterns.ParsePatrialy(allows), nil
}

func selectAllows(d *sqlx.DB, id string) ([]string, error) {
	allows := []string{}
	stmt := "SELECT pattern FROM allows WHERE content_id=$1"
	err := d.Select(&allows, stmt, id)
	if err != nil {
		return nil, err
	}
	return allows, nil
}
