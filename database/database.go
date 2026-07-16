package database

import (
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func InitDatabase() *sql.DB {
	dsn := "host=localhost port=5432 user=postgres password=dikim123 dbname=crud_employee_go sslmode=disable"

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}