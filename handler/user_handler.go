package handler

import (
	"github.com/abinba/codereview/database"
	"github.com/abinba/codereview/middleware"
	"github.com/abinba/codereview/model"
	"github.com/abinba/codereview/repo"
	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserLogin struct {
	Username string `json:"username" example:"johndoe" description:"The username of the user"`
	Password string `json:"password" description:"The password of the user"`
}


type UserSignup struct {
	FirstName string `json:"first_name" example:"john"`
	LastName string `json:"last_name" example:"doe"`
	Username string `json:"username" example:"johndoe" description:"The username of the user"`
	Password string `json:"password" description:"The password of the user"`
}

func validateUser(username, password string) (bool, string) {
	if len(username) < 3 || len(username) > 30 || !govalidator.IsAlphanumeric(username) {
		return false, "Username must be 3-30 characters long and alphanumeric."
	}
	passwordRequirements := map[string]string{
		"uppercase": `[A-Z]`,        // Use at least one uppercase letter
		"lowercase": `[a-z]`,        // Use at least one lowercase letter
		"number":    `[0-9]`,        // Use at least one digit
		"special":   `[^A-Za-z0-9]`, // Use at least one special character
	}
	for key, regexStr := range passwordRequirements {
		matched := govalidator.StringMatches(password, regexStr)
		if !matched {
			return false, "Password must contain at least one " + key
		}
	}
	if len(password) < 8 {
		return false, "Password must be at least 8 characters long."
	}
	return true, ""
}

// CreateUser godoc
// @Summary Create a new user
// @Description create a new user with the provided information
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body User true "User to create" example("{\"username\": \"John Doe\", \"password\": \"nothing\"}")
// @Success 201 {object} model.User
// @Router /api/v1/register/ [post]
func CreateUser(c *fiber.Ctx) error {
	db := database.DB.Db
	if db == nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Database connection not initialized"})
	}

	credentials := new(UserSignup)

	err := c.BodyParser(credentials)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}

	valid, message := validateUser(credentials.Username, credentials.Password)
	if !valid {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": message})
	}

	user := new(model.User)

	err = db.Where("username = ?", credentials.Username).First(&user).Error
	if err == nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "User already exists"})
	}
	
	if err != gorm.ErrRecordNotFound {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Database query error", "data": err})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(credentials.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Failed to hash password", "data": err})
	}

	credentials.Password = string(hashedPassword)

	userRepo := repo.NewUserRepository(db)
	if userRepo == nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Failed to initialize user repository"})
	}

	err = userRepo.CreateUser(credentials.Username, credentials.Password, credentials.FirstName, credentials.LastName)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create user", "data": err})
	}

	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "User has been created"})
}

// Login godoc
// @Summary User login
// @Description login a user by username and password
// @Tags users
// @Accept  json
// @Produce  json
// @Param credentials body User true "Login credentials" example("{\"username\": \"johndoe\", \"password\": \"p@ssword123\"}")
// @Success 200 {string} string "login successful"
// @Failure 401 {string} string "invalid credentials"
// @Router /api/v1/login [post]
func Login(c *fiber.Ctx) error {
	db := database.DB.Db
	credentials := new(UserLogin)

	if err := c.BodyParser(credentials); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Invalid input", "data": err})
	}

	valid, message := validateUser(credentials.Username, credentials.Password)
	if !valid {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": message})
	}

	user := new(model.User)
	err := db.Where("username = ?", credentials.Username).First(&user).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "User not found"})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid credentials"})
	}

	token, err := middleware.GenerateJWT(user.UserID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Failed to generate token", "data": err})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success", "message": "Login successful", "token": token, "username": credentials.Username, "user_id": user.UserID})
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
		FirstName string `json:"first_name"`
		LastName string `json:"last_name"`
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

	user.FirstName = updateUserData.FirstName
	user.LastName = updateUserData.LastName
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

	err := db.Delete(&user, "user_id = ?", id).Error

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete user", "data": nil})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User deleted"})
}
