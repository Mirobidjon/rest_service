package api

import (
	"encoding/json"
	"fmt"
	_ "task/rest_service/api/docs"
	"task/rest_service/api/handlers"
	"task/rest_service/config"

	swagger "github.com/arsmn/fiber-swagger/v2"
	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// SetUpRouter godoc
// @description Hamkorlab Task API
// @termsOfService https://udevs.io
func SetUpRouter(h handlers.Handler, cfg config.Config) *fiber.App {
	r := fiber.New(fiber.Config{JSONEncoder: json.Marshal, BodyLimit: 100 * 1024 * 1024})
	r.Use(logger.New(), cors.New())
	r.Use(MaxAllowed(5000))

	r.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to The Task API")
	})

	r.Get("/swagger/*", swagger.HandlerDefault)

	r.Get("/ping", h.Ping)
	r.Get("/config", h.GetConfig)

	r.Post("/phone", h.CreatePhone)
	r.Get("/phone", h.GetPhoneList)
	r.Get("/phone/:phone_id", h.GetPhoneByID)
	r.Put("/phone", h.UpdatePhone)
	r.Delete("/phone/:phone_id", h.DeletePhone)

	return r
}

func MaxAllowed(n int) fiber.Handler {
	var countReq int64
	sem := make(chan struct{}, n)
	acquire := func() {
		sem <- struct{}{}
		countReq++
		fmt.Println("countRequest: ", countReq)
	}

	release := func() {
		select {
		case <-sem:
		default:
		}
		countReq--
		fmt.Println("countResp: ", countReq)
	}

	return func(c *fiber.Ctx) error {
		acquire()       // before request
		defer release() // after request

		return c.Next()
	}
}
