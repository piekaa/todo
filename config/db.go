package config

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

func openDb() (*sql.DB, error) {
	return sql.Open("mysql", "user:password@/dbname")
}