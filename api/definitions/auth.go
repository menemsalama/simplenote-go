package definitions

// SigninResponse example
type SigninResponse struct {
	UserResponse
	AccessToken string `json:"accessToken"`
}
