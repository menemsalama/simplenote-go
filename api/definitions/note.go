package definitions

import "net/http"

// NoteRequest example
type NoteRequest struct {
	Title string
	Body  string
}

// Validate helper to validate the coming request data
func (note *NoteRequest) Validate() error {
	// TODO: add validation using ozzo
	return nil
}

// Bind for chi render.Renderer helper
func (note *NoteRequest) Bind(r *http.Request) error {
	if err := note.Validate(); err != nil {
		return err
	}

	return nil
}

// NoteResponse example
type NoteResponse struct {
	ID     uint   `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID uint   `json:"-"`
	// gorm attributes are mandatory for querying
	User UserResponse `json:"user" gorm:"foreignkey:UserID;association_foreignkey:ID"`
	GormFields
}

// TableName returns notes table name
// to be compatible with gorm querying interface (gorm.Model)
func (*NoteResponse) TableName() string {
	return "notes"
}
