package handler

import (
	"github.com/abinba/codereview/database"
	"github.com/abinba/codereview/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// GetUserCodeSnippets godoc
// @Summary Get user's code snippets
// @Description Retrieve all code snippets created by a specific user
// @Tags User Code Snippets
// @Accept  json
// @Produce  json
// @Param user_id path string true "User ID"
// @Success 200 {array} model.CodeSnippet
// @Router /api/v1/user_code_snippets/{user_id} [get]
func GetUserCodeSnippets(c *fiber.Ctx) error {
	db := database.DB.Db
	userID := c.Params("id")

	var userUUID uuid.UUID
	if err := userUUID.UnmarshalText([]byte(userID)); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error", "message": "Invalid user ID format", "data": nil,
		})
	}

	var code_snippets []model.CodeSnippet
	db.Where("user_id = ?", userUUID).Preload("CodeSnippetVersions").Preload("User").Find(&code_snippets)

	if len(code_snippets) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status": "error", "message": "No code snippets found for the user", "data": nil,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status": "success", "message": "User code snippets found", "data": code_snippets,
	})
}
