package config

import (
	clowder "github.com/redhatinsights/app-common-go/pkg/api/v1"
)

var AppConfig *clowder.AppConfig

func init() {
	if clowder.IsClowderEnabled() {
		AppConfig = clowder.LoadedConfig
	}
	AppConfig.MetricsPort = 8080
}
