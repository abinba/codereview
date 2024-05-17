package handler

import (
	"fmt"

	"github.com/abinba/codereview/database"
	"github.com/abinba/codereview/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// GetAllReviewComments godoc
// @Summary Get all review comments
// @Description Get all review comments
// @Tags Review Comments
// @Accept  json
// @Produce  json
// @Success 200 {array} model.ReviewComment
// @Router /api/v1/review_comment/ [get]
func GetReviewComments(c *fiber.Ctx) error {
	db := database.DB.Db

	codeSnippetId := c.Params("id")
	
	var review_comments []model.ReviewComment

	// TODO: exclude user.password :)
	db.Preload("User").
		Order("created_at desc").
		Find(&review_comments, "code_snippet_version_id = ?", codeSnippetId)

	if len(review_comments) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status": "error", "message": "Review comments not found", "data": nil,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status": "success", "message": "Review comments found", "data": review_comments,
	})
}


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
	
	// TODO: get the user id from the JWT token, not from the request.
	// Or check if the user id from the JWT token is the same as the user id in the request.
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

	var codeSnippetVersion model.CodeSnippetVersion
	db.Model(&model.CodeSnippetVersion{}).Where("code_snippet_version_id = ?", review_comment.CodeSnippetVersionID).First(&codeSnippetVersion)

	var codeSnippet model.CodeSnippet
	db.Model(&model.CodeSnippet{}).Where("code_snippet_id = ?", codeSnippetVersion.CodeSnippetID).First(&codeSnippet)

	// TODO: use message queue and worker in the future for non-blocking.
	if codeSnippet.UserID != nil && *codeSnippet.UserID != uuid.Nil {
		notification := model.Notification{
			UserID:           *codeSnippet.UserID,
			NotificationType: "CodeReview",
			Text:             "<a href='/code_snippet/" + codeSnippet.CodeSnippetID.String() + "'>Your code has been reviewed! Check it out!</a>",
		}
		if notifResult := db.Create(&notification); notifResult.Error != nil {
			fmt.Println("Error creating notification")
		}
	}

	return c.Status(201).JSON(fiber.Map{
		"status": "success", "message": "Review comment has been created", "data": review_comment,
	})
}
