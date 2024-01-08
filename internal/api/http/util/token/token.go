package token

import (
	"time"

	"github.com/HermanPlay/backend/internal/api/http/constant"
	"github.com/HermanPlay/backend/internal/config"
	"github.com/golang-jwt/jwt"
)

func GenerateToken(user_id uint, cfg *config.Config) (string, error) {
	token_lifespan := constant.TokenHourLifespan

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(token_lifespan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(cfg.App.Api_secret))
}
