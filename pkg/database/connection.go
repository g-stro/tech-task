package database

import (
	"database/sql"
)

type Connection struct {
	DB *sql.DB
}

func NewConnection(driverName, dsn string) (*Connection, error) {
	conn, err := sql.Open(driverName, dsn)
	if err != nil {
		return nil, err
	}
	return &Connection{DB: conn}, nil
}

func (conn *Connection) Close() {
	_ = conn.DB.Close()
	conn.DB = nil
}
