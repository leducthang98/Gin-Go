package rdb

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go-graphql-boilderplate/config"
	"log"
)

type MysqlHandler struct {
	Conn *sqlx.DB
}

var mysqlInstance = &MysqlHandler{}

func Init(cfg *config.Mysql) error {
	connUrl := fmt.Sprintf("%s:%s@(%s:%d)/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
	db, err := sqlx.Connect("mysql", connUrl)
	if err != nil {
		log.Fatalln(err)
	}
	mysqlInstance.Conn = db
	return nil
}

func GetInstance() SQLHandler {
	return mysqlInstance
}

func (m *MysqlHandler) Query(dest interface{}, sql string, args ...interface{}) (err error) {
	err = m.Conn.Select(dest, sql, args...)
	return
}

func (m *MysqlHandler) Close() error {
	if m != nil {
		return m.Conn.Close()
	}

	return nil
}
