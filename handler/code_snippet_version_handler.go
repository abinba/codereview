package handler

import (
	"github.com/abinba/codereview/database"
	"github.com/abinba/codereview/model"
	"github.com/gofiber/fiber/v2"
)

// CreateCodeSnippetVersion godoc
// @Summary Create a code snippet version
// @Description Create a code snippet version
// @Tags Code Snippets Versions
// @Accept  json
// @Produce  json
// @Param code_snippet body model.CodeSnippetVersion true "Code Snippet information to create"
// @Success 201 {object} model.CodeSnippetVersion
// @Router /api/v1/code_snippet_version/ [post]
func CreateCodeSnippetVersion(c *fiber.Ctx) error {
	db := database.DB.Db

	code_snippet := new(model.CodeSnippetVersion)
	if err := c.BodyParser(code_snippet); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error", "message": "Could not parse request", "data": err,
		})
	}

	if result := db.Create(&code_snippet); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error", "message": "Could not create code snippet version", "data": result.Error,
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"status": "success", "message": "Code Snippet Version has been created", "data": code_snippet,
	})
}
