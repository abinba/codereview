package router

import (
	"github.com/abinba/codereview/handler"
	"github.com/abinba/codereview/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Use(middleware.Security)

	api := app.Group("/api")
	v1 := api.Group("/v1")

	// Login and registration
	v1.Post("/register", handler.CreateUser)
	v1.Post("/login", handler.Login)

	// Users
	user := v1.Group("/user")
	user.Get("/:id", middleware.JWTProtected(), handler.GetSingleUser)
	user.Put("/:id", middleware.JWTProtected(), handler.UpdateUser)
	user.Delete("/:id", middleware.JWTProtected(), handler.DeleteUserByID)

	// Code snippet
	code_snippet := v1.Group("/code_snippet")
	code_snippet.Get("/", handler.GetAllCodeSnippets)
	code_snippet.Get("/:id", handler.GetSingleCodeSnippet)
	code_snippet.Post("/", handler.CreateCodeSnippet)

	// Versions of the code snippet
	code_snippet_version := v1.Group("/code_snippet_version")
	code_snippet_version.Post("/", handler.CreateCodeSnippetVersion)

	// User code snippets
	user_code_snippets := v1.Group("/user_code_snippet")
	user_code_snippets.Get("/:id", middleware.JWTProtected(), handler.GetUserCodeSnippets)

	// Review comments
	review_comment := v1.Group("/review_comment")
	review_comment.Post("/", handler.CreateReviewComment)

	// Programming languages
	program_language := v1.Group("/program_language")
	program_language.Get("/", handler.GetAllProgramLanguages)
	program_language.Post("/", handler.CreateProgramLanguage)

	// Notifications
	notifications := v1.Group("/notifications")
	notifications.Get("/:id", middleware.JWTProtected(), handler.GetNotificationsByUserID)
}
