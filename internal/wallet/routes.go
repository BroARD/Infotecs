package wallet

import "github.com/labstack/echo/v4"


func MapRoutes(walletGroup *echo.Group, h Handlers) {
	walletGroup.GET("/:wallet_id/balance", h.GetByID())
}
