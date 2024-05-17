package handler

import (
	"github.com/abinba/codereview/database"
	"github.com/abinba/codereview/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// GetNotificationsByUserID godoc
// @Summary Get user notifications
// @Description Retrieve all notifications for a specific user
// @Tags Notifications
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {array} model.Notification
// @Router /api/v1/notifications/{id} [get]
func GetNotificationsByUserID(c *fiber.Ctx) error {
	db := database.DB.Db
	userID := c.Params("id")

	var userUUID uuid.UUID
	if err := userUUID.UnmarshalText([]byte(userID)); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error", "message": "Invalid user ID format", "data": nil,
		})
	}

	var notifications []model.Notification
	db.Where("user_id = ?", userUUID).Preload("User").Order("created_at desc").Find(&notifications)

	if len(notifications) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status": "error", "message": "No notifications found for the user", "data": nil,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status": "success", "message": "Notifications found", "data": notifications,
	})
}
