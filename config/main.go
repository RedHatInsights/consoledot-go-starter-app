package config

import (
	"fmt"
	"os"
	"strconv"

	clowder "github.com/redhatinsights/app-common-go/pkg/api/v1"

	"github.com/joho/godotenv"
)

const (
	envFile          = "local.env"
	hostEnvVar       = "HOST"
	publicPortEnvVar = "PUBLIC_PORT"
)

func Load() *Config {
	c := Config{
		EnvVars:   make(map[string]string),
		AppConfig: &clowder.AppConfig{},
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
