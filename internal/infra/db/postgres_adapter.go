package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

type postgresAdapter struct {
	db *sql.DB
}

func NewPostgresAdapter() (*postgresAdapter, error) {
	dsn := fmt.Sprintf(
		`host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo`,
		viper.GetString(`db_host`),
		viper.GetString(`db_user`),
		viper.GetString(`db_password`),
		viper.GetString(`db_name`),
		viper.GetString(`db_port`),
	)

	db, err := sql.Open(`postgres`, dsn)
	if err != nil {
		log.Printf("Failed to connect on database: %v\n", err)
		return nil, err
	}

	return &postgresAdapter{
		db: db,
	}, nil
}

func (pa *postgresAdapter) Query(statement string, params ...interface{}) (*sql.Rows, error) {
	rows, err := pa.db.Query(statement, params...)
	return rows, err
}

func (pa *postgresAdapter) Close() {
	pa.db.Close()
}
