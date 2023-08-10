package utils_test

import (
	"fmt"
	"testing"

	"github.com/charlygame/CatGameService/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func TestJwtSignSuccess(t *testing.T) {
	claims := utils.WxGameClaims{
		jwt.RegisteredClaims{
			Issuer:   "charly",
			Subject:  "subject",
			ID:       "123",
			Audience: []string{"somebody_else"},
		},
	}
	token, err := utils.JwtSign(claims)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(token)
	assert.NotEqual(t, "", token)
}

func TestJwtVerifySuccess(t *testing.T) {
	claims := utils.WxGameClaims{
		jwt.RegisteredClaims{
			Issuer:   "charly",
			Subject:  "subject",
			ID:       "123",
			Audience: []string{"somebody_else"},
		},
	}
	token, err := utils.JwtSign(claims)
	if err != nil {
		t.Error(err)
	}

	result, err := utils.JwtVerify(token)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, claims.ID, result.ID)
}
