package utils

import "github.com/golang-jwt/jwt/v5"

type WxGameClaims struct {
	jwt.RegisteredClaims
}

var SIGNING_KEY = []byte("demo-game")

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
