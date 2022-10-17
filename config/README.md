# Configurations
- The package supports the environment configuratons and `.env` file for development purposes

## Example
```
Project struct {
    Name       string `env:"PROJECT_NAME,default=golang_project"`
    LogLevel   string `env:"LOG_LEVEL,default=info"`
    DevMode    bool   `env:"DEV_MODE,default=false"`
    HashCost   int    `env:"HASH_COST,default=10"`
    DefaultOTP string `env:"DEFAULT_OTP,default=123456"`
    OTPLength  int    `env:"OTP_LENGTH,default=6"`
    Version    string `env:"VERSION,default=v1.0.0"`
}
```

If you want to override the `Name` attribute you need to create `.env` file and variable like below
```
PROJECT_NAME='AWESOME_PROJECT'
```
