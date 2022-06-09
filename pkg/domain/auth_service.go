package domain

import (
	"github.com/dgrijalva/jwt-go"
	"net/http"
)

//AuthService Auth service
type AuthService interface {
	ValidateToken(r *http.Request) (*jwt.Token, error)
	TokenValid(r *http.Request) error
	CreateToken(userID string) (*TokenDetails, error)
	CreateAuth(userID string, td *TokenDetails) error
	DeleteTokens(authD *AccessDetails) error
	ExtractTokenMetadata(r *http.Request) (*AccessDetails, error)
	FetchAuth(authD *AccessDetails) (uint64, error)
	DeleteAuth(givenUUID string) (int64, error)
}
