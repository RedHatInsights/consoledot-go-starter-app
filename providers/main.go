package providers

import (
	"github.com/RedHatInsights/consoledot-go-starter-app/config"
	"github.com/RedHatInsights/consoledot-go-starter-app/providers/database"
	"github.com/rs/zerolog/log"
)

type Providers struct {
	DBConnectionPool database.ConnectionPool
	Config           *config.Config
}

// DBProviderGuard returns true if the database connection pool is not nil
// If there is no connection it will attempt to establish one JIT
// If the connection attempt fails it will return false
// If the connection attempt succeeds it will return true
// This guard function should be called before attempting to use the database connection pool
func (p *Providers) DBProviderGuard() bool {
	// We've got a connection pool so we're good
	if p.DBConnectionPool != nil {
		return true
	}
	// Try to establish a connection JIT
	p.DBConnectionPool = dbConnect(p.Config)
	// Return the result after the connection attempt
	return p.DBConnectionPool != nil
}

func Init(conf *config.Config) (Providers, func(Providers)) {
	providers := Providers{
		DBConnectionPool: nil,
	}
	if conf.HasDBProvider() {
		providers.DBConnectionPool = dbConnect(conf)

	}
	return providers, func(pManager Providers) {
		if pManager.DBConnectionPool != nil {
			log.Info().Msg("Closing database connection pool")
			pManager.DBConnectionPool.Close()
		}
	}
}

// dbConnect returns the datbase connection pool
func dbConnect(conf *config.Config) database.ConnectionPool {
	connPool, err := database.Connect(conf)
	if err != nil {
		log.Error().Err(err).Msg("Error connecting to database. ")
	}
	return connPool
}
