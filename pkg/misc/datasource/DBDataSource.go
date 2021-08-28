package datasource

import (
	"database/sql"
)

type DBDataSource interface {
	GetDatabase() *sql.DB
}
