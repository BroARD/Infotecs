package server

import (
	"Infotecs/internal/middleware"
	"Infotecs/internal/transaction"
	"Infotecs/internal/wallet"

	"github.com/labstack/echo/v4"
)

func (s *Server) MapHandlers(e *echo.Echo) error {
	walletRepository := wallet.NewWalletRepository(s.db)
	transRepository := transaction.NewTransRepository(s.db)

	walletService := wallet.NewWalletService(walletRepository, s.logger)
	transService := transaction.NewTransService(transRepository, s.logger, walletRepository)

	walletHandlers := wallet.NewWalletHandlers(walletService, s.logger)
	transHandlers := transaction.NewTransHandlers(transService, s.logger)

	mw := middleware.NewMiddlewareManager(s.cfg, s.logger)

	v1 := e.Group("/api")
	walletGroup := v1.Group("/wallet")
	transGroup := v1.Group("")

	wallet.MapRoutes(walletGroup, walletHandlers)
	transaction.MapRoutes(transGroup, transHandlers, mw)

	routes := e.Routes()
	for _, route := range routes {
		s.logger.Infof("Method: %s, Path: %s\n", route.Method, route.Path)
	}

	return nil
}
