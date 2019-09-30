package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/jinzhu/copier"
	"github.com/menemsalama/simplenote-go/api/definitions"
	"github.com/menemsalama/simplenote-go/internal/auth"
	"github.com/menemsalama/simplenote-go/internal/database"
	"github.com/menemsalama/simplenote-go/models"
)

// CreateUser godoc
// @Tags users
// @Router /users [post]
// @Summary Creates a new user.
// @Description If user name is "exists", error will be returned.
// @Accept json
// @Produce json
// @Param user body definitions.UserRequest true "Add user"
// @Success 200 {object} definitions.UserResponse
func CreateUser(w http.ResponseWriter, r *http.Request) {
	data := &definitions.UserRequest{}
	userModel := &models.User{}
	userResponse := &definitions.UserResponse{}

	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, InvalidRequestError(err))
		return
	}

	copier.Copy(userModel, data)

	if errors := database.PG.Create(userModel).GetErrors(); len(errors) > 0 {
		render.Render(w, r, InvalidRequestError(errors...))
		return
	}

	copier.Copy(userResponse, userModel)

	render.JSON(w, r, userResponse)
}

// ListUsers godoc
// @Tags users
// @Security BearerAuth
// @Router /users [get]
// @Summary List all users.
// @Description returns list of users
// @Produce json
// @Success 200 {array} definitions.UserResponse
func ListUsers(w http.ResponseWriter, r *http.Request) {
	users := &[]definitions.UserResponse{}

	database.PG.Find(users)

	render.JSON(w, r, users)
}

// GetUser godoc
// @Tags users
// @Security BearerAuth
// @Router /users/{id} [get]
// @Summary Get a user.
// @Description returns a user by or ID or current by passing current
// @Produce json
// @Param id path string true "id of a user or current"
// @Success 200 {object} definitions.UserResponse
func GetUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "current" {
		userJWT := auth.FromContext(r.Context())
		if userJWT.HasError() {
			render.Render(w, r, InvalidRequestError(userJWT.Error))
			return
		}

		id = fmt.Sprintf("%v", userJWT.GetClaims("user_id"))
	}

	userResponse := definitions.UserResponse{}

	if errors := database.PG.Where("id = ?", id).First(&userResponse).GetErrors(); len(errors) > 0 {
		render.Render(w, r, InvalidRequestError(errors...))
		return
	}

	render.JSON(w, r, userResponse)
}
