package entity

import (
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthToken struct {
	RefreshToken string `json:"refresh_token"`
	Token        string `json:"access_token"`
}
type User struct {
	Guid interface{}
}
type AuthTokenClaim struct {
	*jwt.StandardClaims
	User User
}
type Tokens struct {
	DB        *mongo.Client
	SecretKey string
}
type TokenHandler struct {
	T *Tokens
}
type Session struct {
	RefreshToken string
	User         User
}
