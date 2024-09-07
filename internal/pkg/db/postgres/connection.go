package postgres

import (
	"database/sql"
	"fmt"
	"github.com/RajabovIlyas/golang-crud/config"
	"github.com/RajabovIlyas/golang-crud/internal/database"
	"log"
)

func NewPsqlDB(cfg *config.Config) (*database.Queries, *sql.DB, error) {

	dataSourceName := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		cfg.Postgres.PostgresqlUser,
		cfg.Postgres.PostgresqlPassword,
		cfg.Postgres.PostgresqlHost,
		cfg.Postgres.PostgresqlPort,
		cfg.Postgres.PostgresqlDBName,
	)

	fmt.Println("data", dataSourceName)

	conn, err := sql.Open("postgres", dataSourceName)

	if err != nil {
		return nil, nil, err
	}

	db := database.New(conn)

	return db, conn, nil

}

func DisconnectPsqlDB(conn *sql.DB) {
	err := conn.Close()
	if err != nil {
		log.Fatal("Error to disconnect postgres: " + err.Error())
	}
}
