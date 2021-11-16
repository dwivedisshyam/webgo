package config

import (
	"os"

	"github.com/dwivedisshyam/webgo/pkg/log"

	"github.com/joho/godotenv"
)

type GoDotEnvProvider struct {
	configFolder string
	logger       log.Logger
}

func NewGoDotEnvProvider(l log.Logger, configFolder string) *GoDotEnvProvider {
	provider := &GoDotEnvProvider{
		configFolder: configFolder,
		logger:       l,
	}

	provider.readConfig(configFolder)

	return provider
}

func (g *GoDotEnvProvider) readConfig(confLocation string) {
	defaultFile := confLocation + "/.env"

	env := os.Getenv("WEBGO_ENV")
	if env == "" {
		env = "local"
	}

	overrideFile := confLocation + "/." + env + ".env"

	err := godotenv.Load(overrideFile)
	if err == nil {
		g.logger.Info("Loaded config from file: ", overrideFile)
	}

	err = godotenv.Load(defaultFile)
	if err == nil {
		g.logger.Info("Loaded config from file: ", defaultFile)
	}
}

func (g *GoDotEnvProvider) Get(key string) string {
	return os.Getenv(key)
}

func (g *GoDotEnvProvider) GetOrDefault(key, defaultValue string) string {
	val := os.Getenv(key)
	if val != "" {
		return val
	}

	return defaultValue
}
