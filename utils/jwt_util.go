package utils

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type WxGameClaims struct {
	jwt.RegisteredClaims
}

var SIGNING_KEY = []byte("demo-game")

func NewWxGameClaims(id string) WxGameClaims {

	nowTime := time.Now()
	expireTime := nowTime.Add(24 * 60 * 60 * time.Second)

	claims := WxGameClaims{
		jwt.RegisteredClaims{
			Issuer:   "charly-game",
			Subject:  "game-demo",
			ID:       id,
			Audience: []string{"somebody_else"},
			IssuedAt: jwt.NewNumericDate(expireTime),
		},
	}
	return claims
}

func JwtSign(claims WxGameClaims) (string, *GameError) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(SIGNING_KEY)
	if err != nil {
		return "", &GameError{StatusCode: 500, Err: err}
	}
	return t, nil
}

func JwtVerify(t string) (*WxGameClaims, *GameError) {
	token, err := jwt.ParseWithClaims(t, &WxGameClaims{}, func(token *jwt.Token) (interface{}, error) {
		return SIGNING_KEY, nil
	})
	if err != nil {
		return nil, &GameError{StatusCode: 500, Err: err}
	}
	claims, ok := token.Claims.(*WxGameClaims)
	if !ok {
		return nil, &GameError{StatusCode: 500, Err: err}
	}
	return claims, nil
}

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Authorization header is required"})
			c.Abort()
			return
		}
		claims, err := JwtVerify(authHeader)
		if err != nil {
			c.JSON(err.StatusCode, gin.H{"message": err.Err.Error()})
			c.Abort()
			return
		}
		c.Set("x-user-id", claims.ID)
		c.Next()
	}
}
