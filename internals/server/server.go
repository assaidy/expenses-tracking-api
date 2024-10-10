package server

import (
	"errors"

	"github.com/assaidy/expenses-tracking-api/internals/storage"
	"github.com/assaidy/expenses-tracking-api/internals/utils"

	pg "github.com/assaidy/expenses-tracking-api/internals/storage/postgres"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type FiberServer struct {
	*fiber.App
	storage storage.Storage
}

func NewFiberServer() *FiberServer {
	// TODO: use a rate limiter
	fs := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "expenses-tracking-api",
			AppName:      "expenses-tracking-api",
			ErrorHandler: errorHandler,
		}),
		storage: pg.NewPgStorage(),
	}
	fs.Use(logger.New())
	return fs
}

func errorHandler(c *fiber.Ctx, err error) error {
	var apiE utils.ApiError
	if errors.As(err, &apiE) {
		c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
		return c.Status(apiE.Code).JSON(apiE)
	}
	code := fiber.StatusInternalServerError
	var fiberE *fiber.Error
	if errors.As(err, &fiberE) {
		code = fiberE.Code
	}
	c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
	return c.Status(code).SendString(err.Error())
}
