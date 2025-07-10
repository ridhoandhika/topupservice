package api

import (
	"topupservice/domain"

	"github.com/gofiber/fiber/v2"
)

type userApi struct {
	userService domain.UserService
}

func User(app *fiber.Group, userService domain.UserService) {
	handler := userApi{
		userService: userService,
	}
	app.Get("user/:id", handler.GetUser)
}

// @Summary Get user by id
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "user id"
// @Router /api/user/{id} [get]
func (a userApi) GetUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	data, err := a.userService.GetUser(ctx.Context(), id)

	if err != nil {
		return ctx.Status(404).JSON(map[string]interface{}{
			"status":  false,
			"message": err.Error(),
			"data":    nil,
		})
	}

	return ctx.Status(200).JSON(data)
}
