package config

import (
	"fmt"
	"os"
	"strings"

	clowder "github.com/redhatinsights/app-common-go/pkg/api/v1"

	"github.com/joho/godotenv"
)

const (
	//Environment Variable file
	envFile = "local.env"
	//CD App Config file used when running locally
	localConfigFile = "local_config.json"
	//Environment Variables
	ginBindAddrEnvVar = "GIN_BIND_ADDR"
	postgres          = "postgres://"
)

func Load() *Config {
	c := Config{}
	c.Load()
	return &c
}

type Config struct {
	AppConfig *clowder.AppConfig
}

func (c *Config) GetApiPath() string {
	appName := "app"
	//iterate through AppConfig.Endpoints looking for the one with the name "app"
	for _, endpoint := range c.AppConfig.Endpoints {
		if endpoint.Name == appName {
			return "/api/" + fmt.Sprintf("%v", endpoint.ApiPath)
		}
	}
	return "/api/" + appName
}

func (c *Config) RouterBindAddress() string {
	host := os.Getenv(ginBindAddrEnvVar) + ":"

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
		c.loadConfigFromLocalFile()
	}
}

func (c *Config) loadConfigFromLocalFile() {
	conf, err := clowder.LoadConfig(localConfigFile)
	if err != nil {
		panic(err)
	}
	c.AppConfig = conf
}

func (c *Config) loadEnvVars() {
	err := godotenv.Load(envFile)
	if err != nil {
		// ... handle error
		//log.Fatalf("Some error occured. Err: %s", err)
		panic(err)
	}
}
