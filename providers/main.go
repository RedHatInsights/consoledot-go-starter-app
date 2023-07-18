package providers

import (
	"os"

	"github.com/RedHatInsights/consoledot-go-starter-app/config"
	"github.com/RedHatInsights/consoledot-go-starter-app/providers/database"
	"github.com/rs/zerolog/log"
)

type Providers struct {
	DBConnectionPool database.ConnectionPool
}

func Init(conf *config.Config) (Providers, func(Providers)) {
	providers := Providers{
		DBConnectionPool: nil,
	}
	if conf.AppConfig.Database != nil {
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
		log.Error().Err(err).Msg("Error connecting to database. Exiting ...")
		os.Exit(1)
	}
	return connPool
}
