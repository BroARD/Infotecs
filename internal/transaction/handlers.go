package transaction

import (
	"Infotecs/internal/entity"
	"Infotecs/internal/transaction/dto"
	"Infotecs/pkg/logging"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Handlers interface {
	Create() echo.HandlerFunc
	GetByCount() echo.HandlerFunc
}

type transHandler struct {
	service Service
	logger  logging.Logger
}

func NewTransHandlers(service Service, logger logging.Logger) Handlers {
	return &transHandler{service: service, logger: logger}
}

func (h *transHandler) Create() echo.HandlerFunc {
	return func (ctx echo.Context) error {
		getTrans := &dto.CreateTransDTO{}
		ctx.Bind(&getTrans)

		resultTransaction := &entity.Transaction{
			ID: uuid.NewString(),
			Status: entity.StatusPending,
			Sender: getTrans.From,
			Receiver: getTrans.To,
			Amount: getTrans.Amount,
		}

		createdTrans, err := h.service.Create(ctx.Request().Context(), resultTransaction)
		if err != nil {
			h.logger.Info("Create: Ошибка при создании транзакции")
			return ctx.JSON(http.StatusBadRequest, err)
		}
		return ctx.JSON(http.StatusCreated, createdTrans)
	}
}

func (h *transHandler) GetByCount() echo.HandlerFunc {
	return func (ctx echo.Context) error{
		trans_count, err := strconv.Atoi(ctx.QueryParam("count"))
		if err != nil || trans_count < 0{
			h.logger.Info("GetByCount: Неправильнйы параметр кол-ва транзакций")
			return ctx.JSON(http.StatusNotFound, "Неверный параметр")
		}

		transList, err := h.service.GetTransactionsByCount(ctx.Request().Context(), trans_count)
		if err != nil {
			h.logger.Info("GetByCount: Ошибка при получении списка транзакций")
			return ctx.JSON(http.StatusNotFound, "Could not get transactions list")
		}

		return ctx.JSON(http.StatusOK, transList)
	}
}