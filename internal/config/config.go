package config

import (
	"fmt"
	"os"

	"github.com/Netflix/go-env"
	"github.com/abdukhashimov/go_api/internal/pkg/logger"
	"github.com/abdukhashimov/go_api/pkg/logger/options"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

type AppMode string

const (
	DEVELOPMENT AppMode = "DEVELOPMENT"
	PRODUCTION  AppMode = "PRODUCTION"
)

type Config struct {
	Logging options.Logging `yaml:"logging"`

	Project struct {
		Name    string `env:"PROJECT_NAME" yaml:"name"`
		Mode    string `env:"APPLICATION_MODE"`
		Version string `env:"APPLICATION_VERSION" yaml:"version"`
	} `yaml:"project"`

	Http struct {
		HTTP_HOST string `env:"HTTP_HOST" yaml:"HTTP_HOST"`
		HTTP_PORT int    `env:"HTTP_PORT" yaml:"HTTP_PORT"`
	} `yaml:"http"`

	MongoDB struct {
		URI string `env:"MONGODB_URI"`
	}
}

func Load() *Config {
	var cfg Config

	err := godotenv.Load()
	if err != nil && !os.IsNotExist(err) {
		logger.Log.Warn(".env file is not found")
	}

	appMode := getAppMode()
	configPath, err := getConfigPath(appMode)
	if err != nil {
		panic(err)
	}

	file, err := os.ReadFile(configPath)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(file, &cfg)
	if err != nil {
		panic(err)
	}

	_, err = env.UnmarshalFromEnviron(&cfg)
	if err != nil {
		panic("unmarshal from environment error")
	}

	if err := validateConfig(&cfg); err != nil {
		panic(err)
	}

	return &cfg
}

func getAppMode() AppMode {
	mode := AppMode(os.Getenv("APPLICATION_MODE"))
	if mode != PRODUCTION {
		mode = DEVELOPMENT
	}

	return mode
}

func getConfigPath(appMode AppMode) (string, error) {
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}
	suffix := "Dev"
	if appMode == PRODUCTION {
		suffix = "Prod"
	}

	return fmt.Sprintf("%s/configs/appConfig%s.yaml", path, suffix), nil
}
