package controller

import (
	"goApp/src/internal/assets/factory/dto"
	"goApp/src/internal/service"

	"github.com/gofiber/fiber/v3"
)

type GradeController struct {
	svc service.GradeService
}

func (c GradeController) CreateGrade(ctx fiber.Ctx) error {
	var body dto.GradeDTO
	if err := ctx.Bind().Body(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}

	if body.StudentID == "" {
		body.StudentID = ctx.Params("id")
	}

	res, err := c.svc.AddGrade(ctx.Context(), body)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusCreated).JSON(res)
}

func (c GradeController) GetGradesBySubject(ctx fiber.Ctx) error {
	studentID := ctx.Params("id")
	subject := ctx.Query("subject")
	if subject == "" {
		return fiber.NewError(fiber.StatusBadRequest, "subject is required")
	}

	res, err := c.svc.ListBySubject(ctx.Context(), studentID, subject)
	if err != nil {
		return err
	}

	return ctx.JSON(res)
}

func (c GradeController) GetAverageBySubject(ctx fiber.Ctx) error {
	studentID := ctx.Params("id")
	subject := ctx.Query("subject")
	if subject == "" {
		return fiber.NewError(fiber.StatusBadRequest, "subject is required")
	}

	avg, err := c.svc.AverageBySubject(ctx.Context(), studentID, subject)
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"studentId": studentID,
		"subject":   subject,
		"average":   avg,
	})
}

func NewGradeController(s service.GradeService) *GradeController {
	return &GradeController{svc: s}
}
