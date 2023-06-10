package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	clowder "github.com/redhatinsights/app-common-go/pkg/api/v1"

	"github.com/joho/godotenv"
)

const (
	envFile          = "local.env"
	hostEnvVar       = "HOST"
	publicPortEnvVar = "PUBLIC_PORT"
	postgres         = "postgres://"

	db_user     = "DB_USER"
	db_password = "DB_PASSWORD"
	db_hostname = "DB_HOSTNAME"
	db_port     = "DB_PORT"
	db_name     = "DB_NAME"
)

func Load() *Config {
	c := Config{
		EnvVars: make(map[string]string),
		AppConfig: &clowder.AppConfig{
			Database: &clowder.DatabaseConfig{},
		},
	}
	c.Load()
	return &c
}

type Config struct {
	AppConfig *clowder.AppConfig
	EnvVars   map[string]string
}

func (c *Config) RouterBindAddress() string {
	host := os.Getenv(hostEnvVar) + ":"

	//Append AppConfig.PublicPort to host
	host += fmt.Sprint(*c.AppConfig.PublicPort)

	return host
}

// Get Database Connection String
func (c *Config) DatabaseConnectionString() string {
	dbConnectionStringParts := []string{
		postgres,
		c.AppConfig.Database.Username, ":",
		c.AppConfig.Database.Password, "@",
		c.AppConfig.Database.Hostname, ":",
		fmt.Sprint(c.AppConfig.Database.Port), "/",
		c.AppConfig.Database.Name}
	return strings.Join(dbConnectionStringParts, "")
}

func (c *Config) Load() {
	c.loadEnvVars()
	if clowder.IsClowderEnabled() {
		c.AppConfig = clowder.LoadedConfig
	} else {
		c.setAppConfigFromEnvVars()
	}
}

func (c *Config) loadEnvVars() {
	err := godotenv.Load(envFile)
	if err != nil {
		// ... handle error
		//log.Fatalf("Some error occured. Err: %s", err)
		panic(err)
	}
}

func (c *Config) setAppConfigFromEnvVars() {
	//Set the AppConfig public port from the publicPort const
	c.AppConfig.PublicPort = c.getEnvVarIntPtr(publicPortEnvVar)
	c.AppConfig.Database.Username = os.Getenv(db_user)
	c.AppConfig.Database.Password = os.Getenv(db_password)
	c.AppConfig.Database.Hostname = os.Getenv(db_hostname)
	c.AppConfig.Database.Port = *c.getEnvVarIntPtr(db_port)
	c.AppConfig.Database.Name = os.Getenv(db_name)
}

func (c *Config) getEnvVarIntPtr(envVar string) *int {
	// string to int
	i, err := strconv.Atoi(os.Getenv(envVar))
	if err != nil {
		// ... handle error
		// log.Fatalf("Some error occured. Err: %s", err)
		panic(err)
	}
	return &i
}
