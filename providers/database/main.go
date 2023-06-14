package database

import (
	"context"
	"fmt"
	"os"

	"github.com/RedHatInsights/consoledot-go-starter-app/config"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

// We use an interface here so that we can mock out the database connection in tests
// This is good, because it makes testing easy. This is bad because every time you extend what you
// want to do with the connection pool you will need to extend this interface.
// For example, if you need to use pgxpool.Pool.Query() you'll need to add its method
// signature to this interface. That said it shouldn't be much more than a copy and paste
// from the api docs here: https://pkg.go.dev/github.com/jackc/pgx/v4/pgxpool#Pool
// Feel free to adopt whatever model makes sense to you. You could abandon the interface
// and use one of the pgx mocking projects out there. This just seemed to be simplest for a starter-app.
// Also, if you do end up using this approach remember you'll need to implement mocks
// for whatever you extend the interface to do.
type ConnectionPool interface {
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Close()
}

func Connect(conf *config.Config) (ConnectionPool, error) {
	connPool, err := pgxpool.Connect(context.Background(), conf.DatabaseConnectionString())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return nil, err
	}
	return connPool, nil
}
