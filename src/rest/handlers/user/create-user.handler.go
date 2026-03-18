package user

import (
	"github.com/gofiber/fiber/v2"
	"goboilerplate.com/src/rest/response"
	"goboilerplate.com/src/usecases"
	"goboilerplate.com/src/usecases/user"
)

type CreateUserHandler struct {
	createUserUseCase user.ICreateUserUseCase
}

func NewCreateUserHandler(createUserUseCase user.ICreateUserUseCase) *CreateUserHandler {
	return &CreateUserHandler{
		createUserUseCase: createUserUseCase,
	}
}

func (h *CreateUserHandler) CreateUser(c *fiber.Ctx) error {
	var req user.CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request body",
		})
	}

	resData, err := h.createUserUseCase.Apply(c.UserContext(), req)
	var res response.BaseResponse[any]
	
	if err != nil {
		switch err {
		case usecases.ErrUserAlreadyExists:
			res = response.Responses[response.ConflictResponse]
		case usecases.ErrCannotCreateUser:
			res = response.Responses[response.InternalServerErrorResponse]
		default:
			res = response.Responses[response.InternalServerErrorResponse]
		}
	} else {
		res = response.Responses[response.SuccessResponse]
		res.Data = resData
	}

	return c.Status(res.HttpStatus).JSON(res)
}