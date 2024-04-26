package handler

import (
	"fmt"

	"github.com/abinba/codereview/database"
	"github.com/abinba/codereview/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// CreateReviewComment creates a review comment.
// @Summary Create a review comment
// @Description Adds a new review comment to the database.
// @Tags Review Comments
// @Accept json
// @Produce json
// @Param review_comment body model.ReviewComment true "Review comment information to create"
// @Success 201 {object} model.ReviewComment
// @Router /api/v1/review_comment/ [post]
func CreateReviewComment(c *fiber.Ctx) error {
	db := database.DB.Db

	review_comment := new(model.ReviewComment)
	if err := c.BodyParser(review_comment); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error", "message": "Could not parse request", "data": err,
		})
	}

	if result := db.Create(&review_comment); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error", "message": "Could not create review comment", "data": result.Error,
		})
	}

	var codeSnippet model.CodeSnippet
	db.Model(&model.CodeSnippet{}).Preload("CodeSnippetVersions", "code_snippet_version_id = ?", review_comment.CodeSnippetVersionID).First(&codeSnippet)
	// TODO: use message queue and worker in the future for non-blocking.
	if codeSnippet.UserID != nil && *codeSnippet.UserID != uuid.Nil {
		notification := model.Notification{
			UserID:           *codeSnippet.UserID,
			NotificationType: "CodeReview",
			Text:             "Your code has been reviewed! Check it out at My snippets",
		}
		if notifResult := db.Create(&notification); notifResult.Error != nil {
			fmt.Println("Error creating notification")
		}
	}

	return c.Status(201).JSON(fiber.Map{
		"status": "success", "message": "Review comment has been created", "data": review_comment,
	})
}
