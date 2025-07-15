package api

import (
	"strings"
	"topupservice/domain"
	"topupservice/dto"
	util "topupservice/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type authApi struct {
	authService domain.AuthService
}

func Auth(app *fiber.Group, userService domain.AuthService, authMid fiber.Handler) {
	handler := authApi{
		authService: userService,
	}
	// Menambahkan anotasi Swagger untuk login
	app.Post("auth/login", handler.GenerateToken)
	// Menambahkan anotasi Swagger untuk refresh
	app.Get("auth/refresh", authMid, handler.ValidateToken)
	// Menambahkan anotasi Swagger untuk refresh
	app.Post("auth/register", handler.Register)
}

// @Summary Generate Token for Authentication
// @Tags auth
// @Accept json
// @Produce json
// @Param body body dto.AuthReq true "Login Credentials" // Mendefinisikan parameter yang dikirimkan dalam request body
// @Success 200 "JWT Token"
// @Failure 400 "Invalid Request"
// @Failure 401   "Authentication Failed"
// @Router /api/auth/login [post]
func (a authApi) GenerateToken(ctx *fiber.Ctx) error {
	var req dto.AuthReq
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(util.GetHttpStatus(domain.ErrAuthFailed))
	}

	token, err := a.authService.Login(ctx.Context(), req)

	if err != nil {
		return ctx.SendStatus(util.GetHttpStatus(domain.ErrAuthFailed))
	}

	return ctx.Status(200).JSON(domain.BaseResp{
		Status:  true,
		Message: "success",
		Data:    token,
	})
}

// @Security BearerAuth
// @Summary Refresh JWT Token
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 "New JWT Token"
// @Failure 401 "Authentication Failed"
// @Router /api/auth/refresh [get]
func (a authApi) ValidateToken(ctx *fiber.Ctx) error {
	authHeader := ctx.Get("Authorization")

	// Memparsing token (memisahkan Bearer dan token)
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return ctx.SendStatus(util.GetHttpStatus(domain.ErrAuthFailed))
	}

	token := parts[1]
	user, err := a.authService.Refresh(ctx.Context(), token)
	if err != nil {
		// Jika token tidak valid atau error lain
		return ctx.SendStatus(util.GetHttpStatus(domain.ErrAuthFailed))
	}

	return ctx.Status(200).JSON(domain.BaseResp{
		Status:  true,
		Message: "success",
		Data:    user,
	})
}

// @Summary Register a new user
// @Tags auth
// @Accept json
// @Produce json
// @Param body body dto.UserRegisterReq true "User Registration Request"
// @Success 200 "Registration Success"
// @Failure 400 "Invalid Request"
// @Failure 409 "User already exists"
// @Router /api/auth/register [post]
func (a authApi) Register(ctx *fiber.Ctx) error {
	var req dto.UserRegisterReq
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"status":  false,
			"message": err.Error(),
			"data":    nil,
		})
	}

	isSuccess, err := a.authService.Register(ctx.Context(), req)

	if err != nil {
		return ctx.Status(util.GetHttpStatus(err)).JSON(map[string]interface{}{
			"status":  isSuccess,
			"message": err.Error(),
			"data":    nil,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(domain.BaseResp{
		Status:  isSuccess,
		Message: "success",
		Data:    nil,
	})
}
