package repository

import (
	"database/sql"
 _ "github.com/go-sql-driver/mysql"
)

type MysqlBookRepository struct {
	db *sql.DB
}

func NewMySqlBookRepository() {
	db, err = sql.Open("mysql","")

	if err != nil {
        log.Fatal(err)
    }

	return &MysqlBookRepository{
		db: db
	}
}
