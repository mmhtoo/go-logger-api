package config

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type DatabaseParameters struct {
	Host string
	Port string
	DBName string
	Timeout time.Duration
}

type DatabaseCredentials struct {
	Username string
	Password string
}

type Database struct {
	Connection *sql.DB
	parameters DatabaseParameters
}

func NewDatabase(
	ctx context.Context,
	parameters DatabaseParameters,
	credentials DatabaseCredentials,
) (*Database, error){
	connectionString := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		parameters.Host,
		parameters.Port,
		parameters.DBName,
		credentials.Username,
		credentials.Password,
	)
	conn, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	if err = conn.Ping(); err != nil {
		log.Fatalf("Ping failed: %s", err)
		return nil, err
	}
	database := &Database{
		parameters: parameters,
		Connection: conn,
	}
	return database, nil
}