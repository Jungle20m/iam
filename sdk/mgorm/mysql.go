package mgorm

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type ConnectionType string

const (
	Mysql    string = "mysql"
	Postgres string = "postgres"
)

const (
	_defaultMaxOpenConnection     = 100
	_defaultMaxIdleConnection     = 10
	_defaultConnectionMaxLifetime = time.Hour
)

type Database struct {
	maxOpenConnection     int
	maxIdleConnection     int
	connectionMaxLifetime time.Duration
	connectionType        string

	Connection *gorm.DB
}

func New(dsn string, opts ...Option) (*Database, error) {
	db := &Database{
		maxOpenConnection:     _defaultMaxOpenConnection,
		maxIdleConnection:     _defaultMaxIdleConnection,
		connectionMaxLifetime: _defaultConnectionMaxLifetime,
		connectionType:        Mysql,
	}

	for _, opt := range opts {
		opt(db)
	}

	switch db.connectionType {
	case Mysql:
		fmt.Println("mysql")
		conn, err := sql.Open("mysql", dsn)
		if err != nil {
			return nil, err
		}
		conn.SetMaxOpenConns(db.maxOpenConnection)
		conn.SetMaxIdleConns(db.maxIdleConnection)
		conn.SetConnMaxLifetime(db.connectionMaxLifetime)

		gormDB, err := gorm.Open(mysql.New(mysql.Config{
			Conn: conn,
		}), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
		if err != nil {
			return nil, err
		}

		db.Connection = gormDB
		return db, nil
	case Postgres:
		fmt.Println("postgres")

		gormDB, err := gorm.Open(postgres.New(postgres.Config{
			DSN: dsn,
		}), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
		if err != nil {
			return nil, err
		}
		db.Connection = gormDB
		return db, nil
	default:
		return nil, fmt.Errorf("mgorm only support mysql or postgres")
	}
}

//anhnv:anhnv!@#456@tcp(1.53.252.177:3306)/healthnet?charset=utf8mb4&parseTime=True&loc=Local
