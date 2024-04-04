package router

import (
	"github.com/abinba/codereview/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	user := v1.Group("/user")

	user.Get("/", handler.GetAllUsers)
	user.Get("/:id", handler.GetSingleUser)
	user.Post("/", handler.CreateUser)
	user.Put("/:id", handler.UpdateUser)
	user.Delete("/:id", handler.DeleteUserByID)

	code_snippet := v1.Group("/code_snippet")

	code_snippet.Get("/", handler.GetAllCodeSnippets)
	code_snippet.Get("/:id", handler.GetSingleCodeSnippet)
	code_snippet.Post("/", handler.CreateCodeSnippet)
	code_snippet.Delete("/:id", handler.DeleteCodeSnippetByID)

	program_language := v1.Group("/program_language")

	program_language.Get("/", handler.GetAllProgramLanguages)
	program_language.Post("/", handler.CreateProgramLanguage)
}
