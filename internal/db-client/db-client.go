package db_client

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type DbClient struct {
	ConnString string
	Db         *sql.DB
}

type Profile struct {
	Name string
}

func NewDbClient(connString string) (*DbClient, error) {

	db, err := sql.Open("mysql", connString)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(time.Minute * 5)
	db.SetMaxOpenConns(10)

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	fmt.Println("Successfully pinged the database")
	return &DbClient{
		ConnString: connString,
		Db:         db,
	}, nil
}
