package member

import (
	"ent-study/application"
	"ent-study/application/dto"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	memberApp application.MemberApplication
	validator *validator.Validate
}

func New(memberApp application.MemberApplication) *handler {
	return &handler{
		memberApp: memberApp,
		validator: validator.New(),
	}
}

func (h handler) GetMember(c *fiber.Ctx) error {
	defer c.Response().Header.SetContentType(fiber.MIMEApplicationJSONCharsetUTF8)

	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	member, err := h.memberApp.GetMember(c.Context(), id)
	if err != nil {
		return err
	}

	return c.Status(200).JSON(fiber.Map{
		"code":    "2000",
		"message": "정상처리",
		"data":    member,
	})
}

func (h handler) JoinMember(c *fiber.Ctx) error {
	defer c.Response().Header.SetContentType(fiber.MIMEApplicationJSONCharsetUTF8)

	body := &dto.JoinMemberRequest{}
	if err := c.BodyParser(body); err != nil {
		return err
	}

	if err := h.validator.Struct(body); err != nil {
		return fiber.ErrBadRequest
	}

	if err := h.memberApp.JoinMember(c.Context(), body); err != nil {
		return err
	}

	return c.Status(200).JSON(fiber.Map{
		"code":    "2000",
		"message": "정상처리",
	})
}
