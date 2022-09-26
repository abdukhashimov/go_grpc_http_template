package config

import (
	"log"

	env "github.com/Netflix/go-env"
	"github.com/joho/godotenv"

	"github.com/abdukhashimov/go_api/pkg/logger"
)

type Environment struct {
	logger.Options

	Project struct {
		Name       string `env:"PROJECT_NAME,default=golang_project"`
		LogLevel   string `env:"LOG_LEVEL,default=info"`
		DevMode    bool   `env:"DEV_MODE,default=false"`
		HashCost   int    `env:"HASH_COST,default=10"`
		DefaultOTP string `env:"DEFAULT_OTP,default=123456"`
		OTPLength  int    `env:"OTP_LENGTH,default=6"`
		Version    string `env:"VERSION,default=v1.0.0"`
	}

	Rest struct {
		HOST string `env:"REST_HOST,default=0.0.0.0"`
		Port int    `env:"REST_PORT,default=8080"`
	}

	Jwt struct {
		AccessSecretKey   string `env:"JWT_ACCESS_SECRET_KEY,default=access_secret"`
		RefreshSecretKey  string `env:"JWT_REFRESH_SECRET_KEY,default=refresh_secret"`
		PasscodeSecretKey string `env:"JWT_PASSCODE_SECRET_KEY,default=passcode_secret"`
		AccessExpire      int    `env:"JWT_ACCESS_EXPIRE,default=3600"`
		RefreshExpire     int    `env:"JWT_REFRESH_EXPIRE,default=86400"`
		PasscodeExpire    int    `env:"JWT_PASSCODE_EXPIRE,default=180"`
	}

	Roles struct {
		Client string `env:"ROLE_CLIENT,default=client"`
		Admin  string `env:"ROLE_ADMIN,default=admin"`
		Editor string `env:"ROLE_EDITOR,default=editor"`
	}
}

func Load() *Environment {
	var (
		cfg Environment
	)

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	_, err = env.UnmarshalFromEnviron(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	return &cfg
}
