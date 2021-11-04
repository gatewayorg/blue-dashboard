package jwt

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

const (
	DefaultIss             = "blue-dashboard"
	DefaultExpiredDuration = 24 * time.Hour
)

var (
	DefaultSignMethod = jwt.SigningMethodHS256
)

type Signer interface {
	GetKey() []byte
	GetSignMethod() jwt.SigningMethod
	Sign(id uint64, username string) (string, error)
	SignWithTime(id uint64, username string, t time.Time) (string, error)
	ParseClaims(tokenString string) (*Claims, error)
	GetExpireDuration() time.Duration
}

func NewSigner(opts ...SignerOption) (*signerImpl, error) {
	options := &SignerOptions{
		Iss:            DefaultIss,
		ExpireDuration: DefaultExpiredDuration,
		SignMethod:     DefaultSignMethod,
	}

	for _, o := range opts {
		o(options)
	}

	return &signerImpl{
		rawKey:         []byte(options.Key),
		iss:            options.Iss,
		expireDuration: options.ExpireDuration,
		signMethod:     options.SignMethod,
	}, nil
}

type SignerOptions struct {
	Key            string
	Iss            string
	ExpireDuration time.Duration
	SignMethod     jwt.SigningMethod
}

type SignerOption func(opt *SignerOptions)

func SignKey(key string) SignerOption {
	return func(opts *SignerOptions) {
		opts.Key = key
	}
}

func Iss(iss string) SignerOption {
	return func(opts *SignerOptions) {
		opts.Iss = iss
	}
}

func ExpireDuration(duration time.Duration) SignerOption {
	return func(opts *SignerOptions) {
		opts.ExpireDuration = duration
	}
}

type signerImpl struct {
	rawKey         []byte
	iss            string
	expireDuration time.Duration
	signMethod     jwt.SigningMethod
}

func (p *signerImpl) GetSignMethod() jwt.SigningMethod {
	return p.signMethod
}

func (p *signerImpl) Sign(id uint64, username string) (string, error) {
	return p.SignWithTime(id, username, time.Now())
}

func (p *signerImpl) SignWithTime(id uint64, username string, t time.Time) (string, error) {
	claim := &Claims{
		ID:       id,
		Username: username,
		RegisteredClaims: &jwt.RegisteredClaims{
			Issuer:    p.iss,
			ExpiresAt: jwt.NewNumericDate(t.Add(p.expireDuration)),
			IssuedAt:  jwt.NewNumericDate(t),
		},
	}

	token := jwt.NewWithClaims(p.signMethod, claim)
	return token.SignedString(p.rawKey)
}

func (p *signerImpl) ParseClaims(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// 必要的验证 RS256
		if token.Method == p.signMethod {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(p.iss, false)
		if !checkIss {
			return nil, errors.New("Invalid issuer.")
		}
		return p.GetKey(), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("Invalid token")
}

func (p *signerImpl) GetKey() []byte {
	return p.rawKey
}

func (p *signerImpl) GetExpireDuration() time.Duration {
	return p.expireDuration
}
