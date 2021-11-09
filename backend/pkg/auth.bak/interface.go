package auth_bak

import "context"

type author interface {
	Authenticate(ctx context.Context, username, passwd string) (ID uint64, err error)
}
