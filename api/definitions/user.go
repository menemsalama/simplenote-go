package definitions

import (
	"net/http"

	"github.com/go-ozzo/ozzo-validation"
)

// UserRequest example
type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Validate helper to validate the coming request data
func (user *UserRequest) Validate() error {
	return validation.ValidateStruct(user,
		validation.Field(&user.Username, validation.Required, validation.Length(5, 50)),
		validation.Field(&user.Password, validation.Required, validation.Length(5, 50)),
	)
}

// Bind for chi render.Renderer helper
func (user *UserRequest) Bind(r *http.Request) error {
	if err := user.Validate(); err != nil {
		return err
	}

	return nil
}

// UserResponse example
type UserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	GormFields
}

// TableName returns notes table name
// to be compatible with gorm querying interface (gorm.Model)
func (user *UserResponse) TableName() string {
	return "users"
}
