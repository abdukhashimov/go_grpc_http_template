package logger

type Options struct {
	ProjectName string `env:"PROJECT_NAME,default=golang_project"`
	Environment string `env:"ENVIRONMENT"`
}
