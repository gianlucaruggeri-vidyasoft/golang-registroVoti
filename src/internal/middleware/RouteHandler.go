package middleware

import (
	"goApp/src/internal/controller"
	"goApp/src/internal/mongo"
	"goApp/src/internal/repository"
	"goApp/src/internal/service"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/requestid"
)

func Init() *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: ErrorHandler,
	})

	app.Use(requestid.New())

	app.Use(logger.New(logger.Config{
		Format:     `"[${time}]${reset} |${status}${reset} | ${latency}${reset} |${method}${reset} |${path}${reset} | IP: ${ip} | req: ${locals:requestid}\n"`,
		TimeFormat: "2006-Jan-02 15:04:05 UTC+1",
		TimeZone:   "UTC",
	}))

	RegisterBaseRoutes(app)
	RegisterServiceRoutes(app)

	log.Info("Registro Voti App Init")

	return app
}

func RegisterBaseRoutes(app *fiber.App) {
	app.Get("/", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{"detail": "found"})
	})

	app.Get("/healthz", func(c fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})
}

func RegisterServiceRoutes(app *fiber.App) {
	api := app.Group("/api/registro")

	studentCtrl := controller.NewStudentController(service.NewStudentService(repository.NewStudentRepository(mongo.SetupMongo())))
	gradeCtrl := controller.NewGradeController(service.NewGradeService(repository.NewGradeRepository(mongo.SetupMongo()), repository.NewStudentRepository(mongo.SetupMongo())))

	api.Post("/studenti", studentCtrl.CreateStudent)
	api.Get("/studenti", studentCtrl.GetAll)
	api.Get("/studenti/:id", studentCtrl.GetOne)
	api.Post("/studenti/:id/voti", gradeCtrl.CreateGrade)
	api.Get("/studenti/:id/voti", gradeCtrl.GetGradesBySubject)
	api.Get("/studenti/:id/voti/media", gradeCtrl.GetAverageBySubject)
}
