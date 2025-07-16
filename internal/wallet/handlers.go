package wallet

import (
	"Infotecs/pkg/logging"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handlers interface {
	GetByID() echo.HandlerFunc
}

type walletHandlers struct {
	service Service
	logger  logging.Logger
}

func NewWalletHandlers(service Service, logger logging.Logger) Handlers {
	return &walletHandlers{service: service, logger: logger}
}

func (h *walletHandlers) GetByID() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		wallet_id := ctx.Param("wallet_id")

		wallet, err := h.service.GetWalletByID(ctx.Request().Context(), wallet_id)
		if err != nil {
			h.logger.Info("GetByID: Ошибка при получении кошелька из Сервиса")
			return ctx.JSON(http.StatusBadRequest, "Кошелёк не найден")
		}
		return ctx.JSON(http.StatusOK, wallet)
	}

}
