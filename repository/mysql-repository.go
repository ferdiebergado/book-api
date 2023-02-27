package repository

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type MysqlRepository struct {
	Repository
}

func NewMySqlRepository() *MysqlRepository {
	mysql, err := sql.Open("mysql", "")

	if err != nil {
		log.Fatal(err)
	}

	return &MysqlRepository{Repository{db: mysql}}
}
