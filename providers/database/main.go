package database

import (
	"context"
	"fmt"
	"os"

	"github.com/RedHatInsights/consoledot-go-starter-app/config"
	"github.com/jackc/pgx/v4/pgxpool"
)

func Connect(conf *config.Config) (*pgxpool.Pool, error) {
	connPool, err := pgxpool.Connect(context.Background(), conf.DatabaseConnectionString())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return nil, err
	}
	return connPool, nil
}
