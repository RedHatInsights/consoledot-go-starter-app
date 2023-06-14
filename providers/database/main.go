package database

import (
	"context"
	"fmt"
	"os"

	"github.com/RedHatInsights/consoledot-go-starter-app/config"
	"github.com/jackc/pgx/v5"
)

func Connect(conf *config.Config) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), conf.DatabaseConnectionString())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return nil, err
	}
	return conn, nil
}
