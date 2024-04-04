package handler

import (
	"github.com/abinba/codereview/database"
	"github.com/abinba/codereview/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type User struct {
	Username string `json:"username" example:"johndoe" description:"The username of the user"`
}

// CreateUser godoc
// @Summary Create a new user
// @Description create a new user with the provided information
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body User true "User to create" example("{\"username\": \"John Doe\"}")
// @Success 201 {object} model.User
// @Router /api/v1/user/ [post]
func CreateUser(c *fiber.Ctx) error {
	db := database.DB.Db
	user := new(model.User)

	err := c.BodyParser(user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}

	err = db.Create(&user).Error // TODO: hash passwords
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create user", "data": err})
	}

	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "User has been created", "data": user})
}

// GetAllUsers godoc
// @Summary Get all users
// @Description get details of all users
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {array} model.User
// @Router /api/v1/user/ [get]
func GetAllUsers(c *fiber.Ctx) error {
	db := database.DB.Db

	var users []model.User
	db.Find(&users)

	if len(users) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Users not found", "data": nil})
	}

	return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "Users Found", "data": users})
}

// GetSingleUser godoc
// @Summary Get single user
// @Description get details of user by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} model.User
// @Router /api/v1/user/{id} [get]
func GetSingleUser(c *fiber.Ctx) error {
	db := database.DB.Db
	id := c.Params("id")

	var user model.User
	db.Find(&user, "user_id = ?", id)

	if user.UserID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found", "data": nil})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User Found", "data": user})
}

// UpdateUser godoc
// @Summary Update an existing user
// @Description update user's information by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Param user body User true "User information to update"
// @Success 200 {object} model.User
// @Router /api/v1/user/{id} [put]
func UpdateUser(c *fiber.Ctx) error {
	type updateUser struct {
		Username string `json:"username"`
	}

	db := database.DB.Db
	var user model.User

	id := c.Params("id")
	db.Find(&user, "user_id = ?", id)

	if user.UserID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found", "data": nil})
	}

	var updateUserData updateUser
	err := c.BodyParser(&updateUserData)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}

	user.Username = updateUserData.Username
	db.Save(&user)

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "users Found", "data": user})
}

// DeleteUserByID godoc
// @Summary Delete a user
// @Description delete a user by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 204 "User deleted"
// @Router /api/v1/user/{id} [delete]
func DeleteUserByID(c *fiber.Ctx) error {
	db := database.DB.Db
	var user model.User
	id := c.Params("id")
	db.Find(&user, "user_id = ?", id)

	if user.UserID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found", "data": nil})
	}

	err := db.Delete(&user, "id = ?", id).Error

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete user", "data": nil})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User deleted"})
}
