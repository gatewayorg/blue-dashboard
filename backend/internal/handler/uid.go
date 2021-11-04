package handler

import (
	"context"
	"github.com/gatewayorg/blue-dashboard/pkg/jwt"
	"google.golang.org/grpc/metadata"
	"strings"
)

type UIDProvider interface {
	UID(ctx context.Context) (uint64, error)
}

func NewUIDProvider() UIDProvider {
	return &metadataUIDProvider{}
}

type metadataUIDProvider struct{}

func (p *metadataUIDProvider) UID(ctx context.Context) (ID uint64, err error) {
	incomingContext, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return 0, ErrInvalidContext
	}
	value, ok := incomingContext["authorization"]
	if !ok {
		return 0, ErrInvalidContext
	}

	authorization := strings.TrimPrefix(value[0], "Bearer ")
	claims, err := jwt.Sign.ParseClaims(authorization)
	if err != nil {
		return 0, ErrAuthorization
	}
	return claims.ID, nil
}
