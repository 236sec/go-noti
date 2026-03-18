package user

import (
	"github.com/gofiber/fiber/v2"
	"goboilerplate.com/src/rest/response"
	"goboilerplate.com/src/usecases"
	"goboilerplate.com/src/usecases/user"
)

type LoginUserHandler struct {
	loginUserUseCase user.ILoginUserUseCase
}

func NewLoginUserHandler(loginUserUseCase user.ILoginUserUseCase) *LoginUserHandler {
	return &LoginUserHandler{
		loginUserUseCase: loginUserUseCase,
	}
}

func (h *LoginUserHandler) LoginUser(c *fiber.Ctx) error {
	var req user.LoginUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request body",
		})
	}

	resData, err := h.loginUserUseCase.Apply(c.UserContext(), req)
	var res response.BaseResponse[any]
	
	if err != nil {
		switch err {
		case usecases.ErrUserNotFound:
			res = response.Responses[response.NotFoundResponse]
		case usecases.ErrInvalidCredentials:
			res = response.Responses[response.UnauthorizedResponse]
		case usecases.ErrUserNotAbleToLogin:
			res = response.Responses[response.ForbiddenResponse]
		default:
			res = response.Responses[response.InternalServerErrorResponse]
		}
	} else {
		res = response.Responses[response.SuccessResponse]
		res.Data = resData
	}

	return c.Status(res.HttpStatus).JSON(res)
}