package handler

import (
	"github.com/abinba/codereview/database"
	"github.com/abinba/codereview/model"
	"github.com/gofiber/fiber/v2"
)

// CreateReviewComment creates a review comment.
// @Summary Create a review comment
// @Description Adds a new review comment to the database.
// @Tags Review Comments
// @Accept json
// @Produce json
// @Param review_comment body ReviewComment true "Review comment information to create"
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

	return c.Status(201).JSON(fiber.Map{
		"status": "success", "message": "Review comment has been created", "data": review_comment,
	})
}
