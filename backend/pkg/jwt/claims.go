package jwt

import "github.com/golang-jwt/jwt/v4"

type Claims struct {
	ID       uint64
	Username string
	*jwt.RegisteredClaims
}

func (c *Claims) VerifyIssuer(iss string) bool {
	return c.Issuer == iss
}
