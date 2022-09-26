package logger

type Options struct {
	ProjectName string `env:"PROJECT_NAME,default=golang_project"`
	DevMode     string `env:"DEV_MODE,default=false"`
	LogLevel    string `env:"LOG_LEVEL,default=info"`
}
