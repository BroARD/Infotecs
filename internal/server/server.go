package server

import (
	"Infotecs/config"
	"Infotecs/pkg/logging"
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Server struct {
	echo   *echo.Echo
	cfg    *config.Config
	db     *gorm.DB
	logger logging.Logger
}

func NewServer(cfg *config.Config, db *gorm.DB, logger logging.Logger) *Server{
	return &Server{echo: echo.New(), cfg: cfg, db: db, logger: logger}
}
func (s *Server) Run() error {
	s.logger.Info("Try server starting")
    server := &http.Server{
        Addr:         ":8080",
        ReadTimeout:  15 * time.Second,
        WriteTimeout: 15 * time.Second,
    }
    if err := s.MapHandlers(s.echo); err != nil {
        return err
    }

    go func() {
        s.logger.Info("Server is listening on PORT: 8080")
        if err := s.echo.StartServer(server); err != nil {
            s.logger.Fatal("Error starting Server: ", err)
        }
    }()
    

    quit := make(chan os.Signal, 1)
    signal.Notify(quit, os.Interrupt)
    <-quit

    s.logger.Info("Shutting down server...")

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    return s.echo.Server.Shutdown(ctx)
}