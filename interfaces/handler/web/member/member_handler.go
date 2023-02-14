package member

import (
	"ent-study/application"
	"ent-study/application/dto"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	memberApp application.MemberApplication
}

func New(memberApp application.MemberApplication) *handler {
	return &handler{
		memberApp: memberApp,
	}
}

func (h handler) GetMember(c *fiber.Ctx) error {
	defer c.Response().Header.SetContentType(fiber.MIMEApplicationJSONCharsetUTF8)
	member, err := h.memberApp.GetMember(c.Context(), 1)
	if err != nil {
		return err
	}

	return c.Status(200).JSON(fiber.Map{
		"code":    "2000",
		"message": "정상처리",
		"data":    member,
	})
}

func (handler) JoinMember(c *fiber.Ctx) error {
	defer c.Response().Header.SetContentType(fiber.MIMEApplicationJSONCharsetUTF8)

	body := &dto.JoinMemberRequest{}
	if err := c.BodyParser(body); err != nil {
		return err
	}

	return c.Status(200).JSON(fiber.Map{
		"code":    "2000",
		"message": "정상처리",
	})
}
