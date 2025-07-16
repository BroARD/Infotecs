package transaction

import (
	"Infotecs/internal/middleware"

	"github.com/labstack/echo/v4"
)

func MapRoutes(transGroup *echo.Group, h Handlers, mw *middleware.MiddlewareManager) {
	transGroup.POST("/send", h.Create(), mw.TransMiddleware)
	transGroup.GET("/transactions", h.GetByCount())
}
