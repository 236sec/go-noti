package user

import (
	"github.com/gofiber/fiber/v2"
	"goboilerplate.com/src/rest/response"
	"goboilerplate.com/src/usecases"
	"goboilerplate.com/src/usecases/user"
)

type GetUserHandler struct {
	getUserUseCase user.IGetUserUseCase
}

func NewGetUserHandler(getUserUseCase user.IGetUserUseCase) *GetUserHandler {
	return &GetUserHandler{
		getUserUseCase: getUserUseCase,
	}
}

func (h *GetUserHandler) GetUser(c *fiber.Ctx) error {
	userID := c.Params("id")
	resData, err := h.getUserUseCase.Apply(c.Context(), userID)
	var res response.BaseResponse[any]
	
	if err != nil {
		switch err {
		case usecases.ErrUserNotFound:
			res = response.Responses[response.NotFoundResponse]
		default:
			res = response.Responses[response.InternalServerErrorResponse]
		}
	} else {
		res = response.Responses[response.SuccessResponse]
		res.Data = resData
	}

	return c.Status(res.HttpStatus).JSON(res)
}