package tokens

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/Feinot/authorization-service/internal/entity"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateToken(db *mongo.Client, secretKey string) *entity.Tokens {
	return &entity.Tokens{
		DB:        db,
		SecretKey: secretKey,
	}
}

func GenerateToken(guid string, t *entity.Tokens) (entity.AuthToken, entity.User, error) {
	var u entity.User
	u.Guid = guid
	expiresAt := time.Now().Add(time.Minute * 1).Unix()

	token := jwt.New(jwt.SigningMethodHS512) // SHA512

	token.Claims = entity.AuthTokenClaim{
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
		User: u,
	}

	tokenString, err := token.SignedString([]byte(t.SecretKey))

	if err != nil {
		return entity.AuthToken{}, u, fmt.Errorf("cannot create signed string token: %v", err)
	}

	refreshToken := uuid.New()

	refreshTokenBase64 := base64.StdEncoding.EncodeToString([]byte(refreshToken.String()))

	return entity.AuthToken{
		Token:        tokenString,
		RefreshToken: refreshTokenBase64,
	}, u, nil
}
