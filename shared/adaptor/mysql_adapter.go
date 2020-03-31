package adaptor

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type config interface {
	GetHost() string
	GetPort() string
	GetUser() string
	GetPassword() string
	GetDbname() string
	GetMaxIdleConns() int
	GetMaxOpenConns() int
	GetConnMaxLifetime() time.Duration
}

func NewMysqlConnection(c config) (*sql.DB, error) {
	source := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", c.GetUser(), c.GetPassword(), c.GetHost(), c.GetPort(), c.GetDbname())
	db, err := sql.Open("mysql", source)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(c.GetMaxIdleConns())
	db.SetMaxOpenConns(c.GetMaxOpenConns())
	db.SetConnMaxLifetime(c.GetConnMaxLifetime())

	return db, nil
}
