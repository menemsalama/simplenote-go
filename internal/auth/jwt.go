package auth

import (
	"context"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
)

// TokenAuth .
var TokenAuth *jwtauth.JWTAuth

// secret key for JWT
var secret = []byte(os.Getenv("SECRET"))

func init() {
	TokenAuth = jwtauth.New("HS256", []byte(secret), nil)
}

// VerifierHandler Seek, verify and validate JWT tokens
func VerifierHandler() func(http.Handler) http.Handler {
	return jwtauth.Verifier(TokenAuth)
}

// AuthenticatorHandler alias to jwtauth.Authenticator
var AuthenticatorHandler = jwtauth.Authenticator

// CreateJWTAccess returns new access token
func CreateJWTAccess(userID uint) (string, error) {
	_, tokenString, err := TokenAuth.Encode(jwt.MapClaims{"user_id": userID, "exp": jwtauth.ExpireIn(5 * time.Minute)})
	return tokenString, err
}

// JWT interface for auth Context values
type JWT struct {
	Error  error
	Claims jwt.MapClaims
	Token  *jwt.Token
}

// HasError error check
func (j *JWT) HasError() bool {
	return j.Error != nil
}

// GetClaims returns claims value
func (j *JWT) GetClaims(key string) interface{} {
	return j.Claims[key]
}

// FromContext returns auth context interface
func FromContext(ctx context.Context) JWT {
	j := JWT{}

	j.Token, j.Claims, j.Error = jwtauth.FromContext(ctx)

	return j
}
