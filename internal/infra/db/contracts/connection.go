package contracts

import "database/sql"

type Connection interface {
	Query(statement string, params ...interface{}) (*sql.Rows, error)
	Close()
}
