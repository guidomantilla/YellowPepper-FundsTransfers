package datasource

import (
	"database/sql"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

type DBDataSource interface {
	GetDatabase() *sql.DB
}

type MysqlDataSource struct {
	database *sql.DB
}

func NewMysqlDataSource(username string, password string, url string) *MysqlDataSource {

	url = strings.Replace(url, ":username", username, 1)
	url = strings.Replace(url, ":password", password, 1)

	database, err := sql.Open("mysql", url)
	if err != nil {
		zap.L().Fatal(err.Error())
	}

	if err = database.Ping(); err != nil {
		zap.L().Fatal(err.Error())
	}

	return &MysqlDataSource{
		database: database,
	}
}

func (mysqlDataSource *MysqlDataSource) GetDatabase() *sql.DB {
	return mysqlDataSource.database
}
