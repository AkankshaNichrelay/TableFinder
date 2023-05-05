package database

import (
	"context"
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

type Config struct {
	Addr     string
	Net      string
	User     string
	Password string
	Database string
}

// Mysql database
type MySQL struct {
	config *Config
	log    *log.Logger
	db     *sql.DB
}

// New creates a new MySQL client instance
func New(log *log.Logger, cfg *Config) (*MySQL, error) {
	// TODO: Put these values into config file
	config := mysql.Config{
		User:   "root",
		Passwd: "root",
		Net:    "tcp",
		Addr:   "docker.for.mac.localhost:3306",
		DBName: "tablefinder",
	}

	db, err := sql.Open("mysql", config.FormatDSN())
	log.Println("sql.Open", "connection string", config.FormatDSN(), "err", err)
	if err != nil {
		return nil, err
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Println("Ping Err", pingErr)
		return nil, pingErr
	}
	log.Println("MySQL Database Connection established.")

	mysql := MySQL{
		log: log,
		db:  db,
	}

	return &mysql, nil
}

// FetchRows takes the query and arguments to return count of rows returned MySQL
func (sql *MySQL) FetchRows(ctx context.Context, tag string, result interface{}, query string, args ...interface{}) (int, error) {
	// TODO: Add logic to fetch Rows
	return 0, nil
}
