package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/gatekeeper")

	
		
	if err !=nil{
			return nil,err
		}

	err = db.Ping()
	if err != nil{
		return nil, err
	}
	return db, nil
}