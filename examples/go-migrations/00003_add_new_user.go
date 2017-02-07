package main

import (
	"database/sql"

	"github.com/jinzhu/gorm"
	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up_00003, Down_00003)
}

func Up_00003(tx *sql.Tx) error {
	gormTx, err := gorm.Open("sqlite3", tx)
	if err != nil {
		return err
	}
	return func(tx *gorm.DB) error {
		return tx.Exec("INSERT INTO users VALUES (3, 'user', '', '');").Error
	}(gormTx)
}

func Down_00003(tx *sql.Tx) error {
	gormTx, err := gorm.Open("sqlite3", tx)
	if err != nil {
		return err
	}
	return func(tx *gorm.DB) error {
		return tx.Exec("DELETE FROM users WHERE id = 3;").Error
	}(gormTx)
}
