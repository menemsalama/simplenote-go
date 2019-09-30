package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/jinzhu/copier"
	"github.com/menemsalama/simplenote-go/api/definitions"
	"github.com/menemsalama/simplenote-go/internal/auth"
	"github.com/menemsalama/simplenote-go/internal/database"
	"github.com/menemsalama/simplenote-go/models"
)

// CreateNote godoc
// @Tags notes
// @Security BearerAuth
// @Router /notes [post]
// @Summary Creates a user note.
// @Description creates a note for the signed in user.
// @Accept json
// @Produce json
// @Param note body definitions.NoteRequest true "Add note"
// @Success 200 {object} definitions.NoteResponse
func CreateNote(w http.ResponseWriter, r *http.Request) {
	userJWT := auth.FromContext(r.Context())
	if userJWT.HasError() {
		render.Render(w, r, InvalidRequestError(userJWT.Error))
		return
	}

	data := &definitions.NoteRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, InvalidRequestError(err))
		return
	}

	userID := uint(userJWT.GetClaims("user_id").(float64))

	noteModel := models.Note{UserID: userID}

	copier.Copy(&noteModel, data)

	if errors := database.PG.Create(&noteModel).GetErrors(); len(errors) > 0 {
		render.Render(w, r, InvalidRequestError(errors...))
		return
	}

	database.PG.First(&noteModel.User)

	userResponse := &definitions.NoteResponse{}

	copier.Copy(userResponse, noteModel)

	render.JSON(w, r, userResponse)
}

// ListNoteshandler godoc
// @Tags notes
// @Security BearerAuth
// @Router /notes [get]
// @Summary List all users notes.
// @Description returns a list of notes for all the users.
// @Produce json
// @Success 200 {array} definitions.NoteResponse
func ListNoteshandler(w http.ResponseWriter, r *http.Request) {
	notes := []definitions.NoteResponse{}

	database.PG.
		Preload("User").
		Find(&notes)

	render.JSON(w, r, notes)
}

// UpdateNotehandler godoc
// @Tags notes
// @Security BearerAuth
// @Router /notes/{id} [post]
// @Summary Update a note.
// @Description Updates a note for a user.
// @Accept json
// @Produce json
// @Param id path string true "Note ID"
// @Param note body definitions.NoteRequest true "Note fields"
// @Success 200 {object} definitions.NoteResponse
func UpdateNotehandler(w http.ResponseWriter, r *http.Request) {
	data := &definitions.NoteRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, InvalidRequestError(err))
		return
	}

	updates := map[string]interface{}{"title": data.Title, "body": data.Body}

	noteID, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)
	if err != nil {
		// err = errors.New("Id is not valid")
		render.Render(w, r, InvalidRequestError(err))
	}

	noteModel := &models.Note{UserID: uint(noteID)}

	if database.PG.Preload("User").First(noteModel).RecordNotFound() {
		err := errors.New("Record Not Found")
		render.Render(w, r, InvalidRequestError(err))
	}

	if errors := database.PG.Model(noteModel).Updates(updates).GetErrors(); len(errors) > 0 {
		render.Render(w, r, InvalidRequestError(errors...))
		return
	}

	userResponse := &definitions.NoteResponse{}

	copier.Copy(userResponse, noteModel)

	render.JSON(w, r, userResponse)
}
