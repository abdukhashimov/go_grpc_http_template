package logger

type Options struct {
	ProjectName    string `env:"PROJECT_NAME,default=golang_project"`
	DevMode        bool   `env:"DEV_MODE,default=false"`
	LogLevel       string `env:"LOG_LEVEL,default=info"`
	DateTimeFormat string `env:"DATE_TIME_FORMAT,default:2006-01-02 15:04:05"`
	DateFormat     string `env:"DATE_FORMAT,default:2006-01-02"`
}
