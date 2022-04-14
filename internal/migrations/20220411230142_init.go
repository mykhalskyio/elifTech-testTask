package migrations

import (
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upInit, downInit)
}

func upInit(tx *sql.Tx) error {
	_, err := tx.Exec(`
	CREATE TABLE IF NOT EXISTS banks (
		id SERIAL PRIMARY KEY,
		bank_name VARCHAR(250),
		interest_rate INTEGER,
		max_loan INTEGER,
		min_down_payment INTEGER,
		loan_term NUMERIC(5, 3));
	CREATE TABLE IF NOT EXISTS mortgages (
		id SERIAL PRIMARY KEY,
		initial_loan INTEGER,
		down_payment INTEGER,
		monthly_payment NUMERIC(20, 3),
		bank_id INTEGER REFERENCES banks (id));
	`)
	if err != nil {
		return err
	}
	return nil
}

func downInit(tx *sql.Tx) error {
	_, err := tx.Exec(`
	DROP TABLE IF EXISTS banks CASCADE;
	DROP TABLE IF EXISTS mortgages CASCADE;
	`)
	if err != nil {
		return err
	}
	return nil
}
