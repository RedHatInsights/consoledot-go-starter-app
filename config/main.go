package config

import (
	"fmt"
	"strings"

	clowder "github.com/redhatinsights/app-common-go/pkg/api/v1"

	"github.com/joho/godotenv"
)

const (
	envFile         = "local.env"
	localConfigFile = "cdappconfig.json"
	//This should match the name of the deployment in the cdappconfig.json and your clowdapp
	//We need this to resolve the api path when running on a cluster
	deploymentName = "starter-app-depolyment"
	//Strings
	postgres  = "postgres://"
	apiPrefix = "/api/"
	//Fallbacks
	ginBindAddr = "0.0.0.0"
)

func Load() *Config {
	c := Config{}
	c.Load()
	return &c
}

type Config struct {
	AppConfig *clowder.AppConfig
}

// Get metrics port and prepend with a colon
func (c *Config) GetMetricsPort() string {
	return fmt.Sprintf(":%v", c.AppConfig.MetricsPort)
}

// Get the api path from the deployment name
func (c *Config) GetApiPath() string {
	deploymentEndpoint, err := c.GetDeploymentEndpoint()
	if err != nil {
		return apiPrefix + deploymentName
	}
	return fmt.Sprintf("%v", deploymentEndpoint.ApiPath)
}

// Finds the endpoint in the cdappconfig that matches the deployment name
func (c *Config) GetDeploymentEndpoint() (clowder.DependencyEndpoint, error) {
	for _, endpoint := range c.AppConfig.Endpoints {
		if endpoint.Name == deploymentName {
			return endpoint, nil
		}
	}
	return clowder.DependencyEndpoint{}, fmt.Errorf("no endpoint found for %s", deploymentName)
}

// Get the host to bind to
func (c *Config) GetBindHost() string {
	return ginBindAddr
}

// Get the host to bind to and append the public port
func (c *Config) RouterBindAddress() string {
	host := c.GetBindHost()
	//Append AppConfig.PublicPort to host
	host += fmt.Sprintf(":%v", *c.AppConfig.PublicPort)
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

func (c *Config) HasDBProvider() bool {
	return c.AppConfig.Database != nil
}
