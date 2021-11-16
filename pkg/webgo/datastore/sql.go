package datastore

import (
	"database/sql"

	"github.com/dwivedisshyam/webgo/pkg/log"

	_ "github.com/mattn/go-sqlite3"
)

type SQL struct {
	DB *sql.DB
}

type SQLConfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
	Dialect  string
}

func NewSQL(l log.Logger, c *SQLConfig) *SQL {
	s := &SQL{}

	switch c.Dialect {
	case "sqlite":
		s.DB = sqliteConn(l, c)
		return s
	case "mysql":
		return s
	}

	return s
}

func sqliteConn(l log.Logger, c *SQLConfig) *sql.DB {
	db, err := sql.Open("sqlite3", c.Name)
	if err != nil {
		l.Errorf("could not connect to DB, HostName: %s, Port: %s, Dialect: %s, error: %v\n",
			c.Host, c.Port, c.Dialect, err)
		return nil
	}

	l.Infof("connected to %s %s with Host: %s Port: %s", c.Dialect, c.Name, c.Host, c.Port)

	return db
}

func (s *SQL) Exec(query string, args ...interface{}) (sql.Result, error) {
	return s.DB.Exec(query, args...)
}

func (s *SQL) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return s.DB.Query(query, args...)
}

func (s *SQL) QueryRow(query string, args ...interface{}) *sql.Row {
	return s.DB.QueryRow(query, args...)
}
