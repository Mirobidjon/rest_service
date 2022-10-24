package handlers

import (
	"strconv"
	"task/rest_service/api/http"
	"task/rest_service/config"
	"task/rest_service/storage"

	"github.com/saidamir98/udevs_pkg/logger"

	fiber "github.com/gofiber/fiber/v2"
)

type Handler struct {
	cfg   config.Config
	log   logger.LoggerI
	store storage.StorageI
}

func NewHandler(cfg config.Config, log logger.LoggerI, store storage.StorageI) Handler {
	return Handler{
		cfg:   cfg,
		log:   log,
		store: store,
	}
}

func (h *Handler) handleResponse(c *fiber.Ctx, status http.Status, data interface{}) error {
	switch code := status.Code; {
	case code < 300:
		h.log.Info(
			"response",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
			logger.Any("data", data),
		)
	case code < 400:
		h.log.Warn(
			"response",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
			logger.Any("data", data),
		)
	default:
		h.log.Error(
			"response",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
			logger.Any("data", data),
		)
	}

	return c.Status(status.Code).JSON(http.Response{
		Status:      status.Status,
		Description: status.Description,
		Data:        data,
	})
}

func (h *Handler) getOffsetParam(c *fiber.Ctx) (offset int, err error) {
	offsetStr := c.Query("offset", h.cfg.DefaultOffset)
	return strconv.Atoi(offsetStr)
}

func (h *Handler) getLimitParam(c *fiber.Ctx) (offset int, err error) {
	offsetStr := c.Query("limit", h.cfg.DefaultLimit)
	return strconv.Atoi(offsetStr)
}
