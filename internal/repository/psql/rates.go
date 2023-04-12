package psql

import "database/sql"

type Rates struct {
	db *sql.DB
}

func NewRates(db *sql.DB) *Rates {
	return &Rates{db: db}
}
func CreateRates() {

}
