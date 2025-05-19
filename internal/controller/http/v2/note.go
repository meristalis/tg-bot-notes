package v2

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/meristalis/tg-bot-notes/internal/entity"
	"github.com/meristalis/tg-bot-notes/internal/usecase"
	"github.com/meristalis/tg-bot-notes/pkg/httpserver/handler"
	"github.com/meristalis/tg-bot-notes/pkg/logger"
)

type noteRoutes struct {
	n usecase.Note
	l logger.Interface
	v *validator.Validate
}

func NewNoteRoutes(apiGroup fiber.Router, n usecase.Note, l logger.Interface) {
	r := &noteRoutes{n, l, validator.New(validator.WithRequiredStructEnabled())}

	noteGroup := apiGroup.Group("/notes")
	{
		noteGroup.Get("/", r.getAllNotes)
		noteGroup.Post("/", r.addNote)
	}
}

type getAllNotesResponse struct {
	Notes []entity.Note `json:"notes"`
}

// @Summary     Get all notes
// @Description Get all notes for the user
// @ID          get-all-notes
// @Tags  	    notes
// @Accept      json
// @Produce     json
// @Success     200 {object} getAllNotesResponse
// @Failure     500 {object} handler.Response
// @Router      /v1/notes [get]
func (r *noteRoutes) getAllNotes(ctx *fiber.Ctx) error {
	notes, err := r.n.GetAllNotes(ctx.UserContext())
	if err != nil {
		r.l.Error(err, "http - v1 - getAllNotes")

		return handler.ErrorResponse(ctx, http.StatusInternalServerError, "database problems")
	}

	return ctx.Status(http.StatusOK).JSON(getAllNotesResponse{notes})
}

type addNoteRequest struct {
	Title   string `json:"title"   validate:"required" example:"My First Note"`
	Content string `json:"content" validate:"required" example:"This is the content of my note."`
}

// @Summary     Add a new note
// @Description Add a new note for the user
// @ID          add-note
// @Tags  	    notes
// @Accept      json
// @Produce     json
// @Param       request body addNoteRequest true "New note"
// @Success     200 {object} entity.Note
// @Failure     400 {object} handler.Response
// @Failure     500 {object} handler.Response
// @Router      /v1/notes [post]
func (r *noteRoutes) addNote(ctx *fiber.Ctx) error {
	var request addNoteRequest

	if err := ctx.BodyParser(&request); err != nil {
		r.l.Error(err, "http - v1 - addNote")

		return handler.ErrorResponse(ctx, http.StatusBadRequest, "invalid request body")
	}

	if err := r.v.Struct(request); err != nil {
		r.l.Error(err, "http - v1 - addNote")

		return handler.ErrorResponse(ctx, http.StatusBadRequest, "invalid request body")
	}
	userIDStr := ctx.Locals("user_id").(string)
	userUUID, err := uuid.Parse(userIDStr)
	if err != nil {
		r.l.Error(err, "Cannot parse user_id")
		return handler.ErrorResponse(ctx, http.StatusUnauthorized, "Invalid user ID")
	}
	// Создание новой заметки
	note := entity.Note{
		Title:   request.Title,
		Content: request.Content,
		UserID:  userUUID,
	}

	newNote, err := r.n.AddNote(ctx.UserContext(), note)
	if err != nil {
		r.l.Error(err, "http - v1 - addNote")

		return handler.ErrorResponse(ctx, http.StatusInternalServerError, "failed to add note")
	}

	return ctx.Status(http.StatusCreated).JSON(newNote)
}
