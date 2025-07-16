package middleware

import (
	"Infotecs/config"
	"Infotecs/pkg/logging"
)

type MiddlewareManager struct {
	cfg    *config.Config
	logger logging.Logger
}

func NewMiddlewareManager(cfg *config.Config, logger logging.Logger) *MiddlewareManager {
	return &MiddlewareManager{cfg: cfg, logger: logger}
}
