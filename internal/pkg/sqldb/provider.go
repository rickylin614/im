package sqldb

import (
	"go.uber.org/dig"

	"im/internal/pkg/config"
	"im/internal/pkg/logger"
)

type digIn struct {
	dig.In

	Config *config.Config
	Log    *logger.Logger
}

func NewDB(in digIn) Client {
	return newDB(in)
}
