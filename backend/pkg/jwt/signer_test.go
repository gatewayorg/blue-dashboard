package jwt

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var testSigningKey = []byte("AllYourBase")

type myClaims struct {
	username string
	*jwt.RegisteredClaims
}

func TestJwtSing(t *testing.T) {
	// Create the Claims
	claims := &myClaims{
		username: "jmz",
		RegisteredClaims: &jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(testSigningKey)

	assert.NoError(t, err)
	t.Log(ss)
}

func TestJwtParse(t *testing.T) {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ0ZXN0IiwiZXhwIjoxNjM1NzUzMTgyfQ.POrtrW3vPzVozlx2C6Ww_BGVuFF8HYM16tcH9TzHxPA"
	token, err := jwt.ParseWithClaims(tokenString, &myClaims{}, func(token *jwt.Token) (interface{}, error) {
		return testSigningKey, nil
	})
	assert.NoError(t, err)

	t.Log(token)
	if claims, ok := token.Claims.(*myClaims); ok && token.Valid {
		t.Log(claims)
	}
}
