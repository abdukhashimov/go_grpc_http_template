package logrus

import (
	"github.com/abdukhashimov/go_api/pkg/logger"
	"github.com/abdukhashimov/go_api/pkg/logger/config"
)

// Factory is the receiver for logrus factory
type Factory struct{}

// Build logrus logger
func (_ *Factory) Build(cfg *config.Logging) (logger.Logger, error) {
	l, err := RegisterLogrusLog(cfg)
	if err != nil {
		return nil, err
	}

	return l, nil
}
