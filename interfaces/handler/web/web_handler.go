package web

import (
	app "ent-study/application"
	"ent-study/errs"
	memberHandler "ent-study/interfaces/handler/web/member"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/pkg/errors"
)

type webHandler struct {
	*fiber.App
}

func NewWebHandler() *webHandler {
	return &webHandler{fiber.New(
		fiber.Config{
			ErrorHandler: errorHandler,
		},
	)}
}

func (w *webHandler) Run(memberApp app.MemberApplication) {
	memberHandler := memberHandler.New(memberApp)
	w.Get("/member/:id", memberHandler.GetMember)
	w.Post("/member", memberHandler.JoinMember)

	w.Use(recover.New(recover.Config{
		Next: func(c *fiber.Ctx) bool {
			return true
		},
	}))

	w.Listen(":8080")
}

func errorHandler(c *fiber.Ctx, err error) error {
	defer c.Response().Header.SetContentType(fiber.MIMEApplicationJSONCharsetUTF8)

	var e *fiber.Error
	// fiber 자체 오류일 경우
	if errors.As(err, &e) {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"code":    e.Code,
			"message": e.Message,
		})
	}

	// 서버에서 지정한 오류일 경우
	var ce *errs.Error
	if errors.As(err, &ce) {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"code":    ce.Code(),
			"message": ce.Message(),
		})
	}

	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"code":    fiber.StatusInternalServerError,
		"message": err.Error(),
	})
}
