package handlers

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/jinzhu/copier"
	"github.com/menemsalama/simplenote-go/api/definitions"
	"github.com/menemsalama/simplenote-go/internal/auth"
	"github.com/menemsalama/simplenote-go/internal/database"
	"github.com/menemsalama/simplenote-go/models"
)

// Signin godoc
// @Tags auth
// @Router /auth/signin [post]
// @Summary Creates a new access.
// @Description If user is "exists" and the credentials are valid, user and AccessToken will be returned.
// @Accept json
// @Produce json
// @Param user body definitions.UserRequest true "Add account" Userdasd
// @Success 200 {object} definitions.SigninResponse
func Signin(w http.ResponseWriter, r *http.Request) {
	data := &definitions.UserRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, InvalidRequestError(err))
		return
	}

	user := &models.User{}
	query := database.PG.Where("username = ?", data.Username)
	if errors := query.First(&user).GetErrors(); len(errors) > 0 {
		render.Render(w, r, InvalidRequestError(errors...))
		return
	}

	if user.Password != data.Password {
		render.Render(w, r, UnauthorizedError("Password or Username is incorrect."))
		return
	}

	token, err := auth.CreateJWTAccess(user.ID)
	if err != nil {
		render.Render(w, r, UnauthorizedError("", err))
		return
	}

	userResponse := &definitions.UserResponse{}
	copier.Copy(userResponse, user)

	signinResponse := definitions.SigninResponse{UserResponse: *userResponse, AccessToken: token}
	render.JSON(w, r, signinResponse)
}
