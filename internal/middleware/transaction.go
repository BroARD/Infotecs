package middleware

import (
	"Infotecs/internal/transaction/dto"
	"bytes"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (mw *MiddlewareManager) TransMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		bodyBytes, _ := io.ReadAll(ctx.Request().Body)
		ctx.Request().Body = io.NopCloser(bytes.NewReader(bodyBytes))

		getTrans := &dto.CreateTransDTO{}
		ctx.Bind(&getTrans)
		if getTrans.From == getTrans.To {
			mw.logger.Info("TransMiddleware: Ошибка FromID == ToID")
			return echo.NewHTTPError(http.StatusBadRequest, "Отправитель и Получатель не могут быть одинаковыми")
		}
		if getTrans.Amount < 0 {
			mw.logger.Info("TransMiddleware: Ошибка Amount < 0")
			return echo.NewHTTPError(http.StatusBadRequest, "Amount должен быть положительным")
		}

		ctx.Request().Body = io.NopCloser(bytes.NewReader(bodyBytes))
		return next(ctx)
	}
}
