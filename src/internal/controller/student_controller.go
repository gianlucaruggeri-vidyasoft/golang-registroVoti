package controller

import (
	"github.com/gofiber/fiber/v3"
	"goApp/src/internal/assets/factory/dto"
	"goApp/src/internal/service"
)

type StudentController struct {
	svc service.StudentService
}

func (c *StudentController) CreateStudent(ctx fiber.Ctx) error {
	var body dto.StudentDTO
	if err := ctx.Bind().Body(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}

	res, err := c.svc.Create(ctx.Context(), body)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusCreated).JSON(res)
}

func (c *StudentController) GetAll(ctx fiber.Ctx) error {
	res, err := c.svc.List(ctx.Context())
	if err != nil {
		return err
	}
	return ctx.JSON(res)
}

func (c *StudentController) GetOne(ctx fiber.Ctx) error {
	id := ctx.Params("id")
	res, err := c.svc.Get(ctx.Context(), id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "not found")
	}
	return ctx.JSON(res)
}


func NewStudentController(s service.StudentService) *StudentController {
	return &StudentController{svc: s}
}