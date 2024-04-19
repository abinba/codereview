package handler

import (
	"github.com/abinba/codereview/database"
	"github.com/abinba/codereview/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type CodeSnippetForm struct {
	ProgramLanguageID uuid.UUID `json:"program_language_id" description:"The UUID of the program language"`
	Text              string    `json:"text" description:"The text of the code snippet, its content"`
	IsPrivate         bool      `json:"is_private" description:"Whether the code snippet is private"`
	IsArchived        bool      `json:"is_archived" description:"Whether the code snippet is archived"`
	IsDraft           bool      `json:"is_draft" description:"Whether the code snippet is a draft"`
}

type ProgramLanguageForm struct {
	Name string `json:"name" description:"The name of the program language"`
}

func GenerateRandomUsername(length int) string {
	return uuid.NewString()[:length]
}

// GetAllCodeSnippets godoc
// @Summary Get all code snippets
// @Description Get all code snippets
// @Tags Code Snippets
// @Accept  json
// @Produce  json
// @Success 200 {array} model.CodeSnippet
// @Router /api/v1/code_snippet/ [get]
func GetAllCodeSnippets(c *fiber.Ctx) error {
	db := database.DB.Db

	var code_snippets []model.CodeSnippet
	db.Find(&code_snippets).Order("created_at desc")

	if len(code_snippets) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status": "error", "message": "Code Snippets not found", "data": nil,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status": "success", "message": "Code Snippets Found", "data": code_snippets,
	})
}

// GetSingleCodeSnippet godoc
// @Summary Get a single code snippet
// @Description Get a single code snippet by ID
// @Tags Code Snippets
// @Accept  json
// @Produce  json
// @Param id path int true "Code Snippet ID"
// @Success 200 {object} model.CodeSnippet
// @Router /api/v1/code_snippet/{id} [get]
func GetSingleCodeSnippet(c *fiber.Ctx) error {
	db := database.DB.Db
	id := c.Params("id")

	var code_snippet model.CodeSnippet
	db.Where("code_snippet_id = ?", id).First(&code_snippet)

	if code_snippet.CodeSnippetID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{
			"status": "error", "message": "Code Snippet not found", "data": nil,
		})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Code Snippet found", "data": code_snippet})
}

// CreateCodeSnippet godoc
// @Summary Create a code snippet
// @Description Create a code snippet
// @Tags Code Snippets
// @Accept  json
// @Produce  json
// @Param code_snippet body CodeSnippetForm true "Code Snippet information to create"
// @Success 201 {object} model.CodeSnippet
// @Router /api/v1/code_snippet/ [post]
func CreateCodeSnippet(c *fiber.Ctx) error {
	db := database.DB.Db

	user := new(model.User)
	user.Username = GenerateRandomUsername(8)
	if result := db.Create(&user); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error", "message": "Could not create anonymous user", "data": result.Error,
		})
	}

	code_snippet := new(model.CodeSnippet)
	if err := c.BodyParser(code_snippet); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error", "message": "Could not parse request", "data": nil,
		})
	}

	program_language := new(model.ProgramLanguage)
	db.First(&program_language, code_snippet.ProgramLanguageID)

	code_snippet.UserID = user.UserID
	code_snippet.User = *user
	code_snippet.ProgramLanguage = *program_language
	if result := db.Create(&code_snippet); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error", "message": "Could not create code snippet", "data": result.Error,
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"status": "success", "message": "Code Snippet has been created", "data": code_snippet,
	})
}

// DeleteCodeSnippetByID godoc
// @Summary Delete a code snippet
// @Description Delete a code snippet by ID
// @Tags Code Snippets
// @Accept  json
// @Produce  json
// @Param id path int true "Code Snippet ID"
// @Success 200 {object} model.CodeSnippet
// @Router /api/v1/code_snippet/{id} [delete]
func DeleteCodeSnippetByID(c *fiber.Ctx) error {
	db := database.DB.Db
	id := c.Params("id")

	var code_snippet model.CodeSnippet
	db.Where("code_snippet_id = ?", id).First(&code_snippet)

	if code_snippet.CodeSnippetID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{
			"status": "error", "message": "Code Snippet not found", "data": nil,
		})
	}

	if result := db.Delete(&code_snippet); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error", "message": "Could not delete code snippet", "data": result.Error,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status": "success", "message": "Code Snippet has been deleted", "data": nil,
	})
}

// GetAllProgramLanguages godoc
// @Summary Get all program languages
// @Description Get all program languages
// @Tags Program Languages
// @Accept  json
// @Produce  json
// @Success 200 {array} model.ProgramLanguage
// @Router /api/v1/program_language/ [get]
func GetAllProgramLanguages(c *fiber.Ctx) error {
	db := database.DB.Db

	var program_languages []model.ProgramLanguage
	db.Find(&program_languages)

	if len(program_languages) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status": "error", "message": "Program Languages not found", "data": nil,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status": "success", "message": "Program Languages Found", "data": program_languages,
	})
}

// CreateProgramLanguage godoc
// @Summary Create a program language
// @Description Create a program language
// @Tags Program Languages
// @Accept  json
// @Produce  json
// @Param name body ProgramLanguageForm true "Name of the program language"
// @Success 201 {object} model.ProgramLanguage
// @Router /api/v1/program_language/ [post]
func CreateProgramLanguage(c *fiber.Ctx) error {
	db := database.DB.Db

	program_language := new(model.ProgramLanguage)
	if err := c.BodyParser(program_language); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error", "message": "Could not parse request", "data": nil,
		})
	}

	if result := db.Create(&program_language); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error", "message": "Could not create program language", "data": result.Error,
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"status": "success", "message": "Program Language has been created", "data": program_language,
	})
}
