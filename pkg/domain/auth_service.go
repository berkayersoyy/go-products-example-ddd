package domain

import (
	"github.com/dgrijalva/jwt-go"
	"net/http"
)

type AuthService interface {
	ValidateToken(r *http.Request) (*jwt.Token, error)
	TokenValid(r *http.Request) error
	CreateToken(userid uint) (*TokenDetails, error)
	CreateAuth(userid uint, td *TokenDetails) error
	DeleteTokens(authD *AccessDetails) error
	ExtractTokenMetadata(r *http.Request) (*AccessDetails, error)
	FetchAuth(authD *AccessDetails) (uint64, error)
	DeleteAuth(givenUuid string) (int64, error)
}
