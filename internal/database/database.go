package database

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

// type Config struct {
// 	Addr     string
// 	Net      string
// 	User     string
// 	Password string
// 	Database string
// }

type Mysql struct {
	// config *Config
	db *sql.DB
}

func New() (*Mysql, error) {
	fmt.Println("Inside database New")
	config := mysql.Config{
		User:   "root",
		Passwd: "root",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "test",
	}

	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		return nil, err
	}

	pingErr := db.Ping()
	if pingErr != nil {
		fmt.Println("Ping Err", pingErr)
		return nil, pingErr
	}
	fmt.Println("Connected!")

	mysql := Mysql{
		db: db,
	}

	return &mysql, nil
}

// func (ms *Mssql) connect() error {
// 	connParams := url.Values{}
// 	connParams.Add("")
// }
