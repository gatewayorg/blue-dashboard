package request

import (
	"errors"
)

var (
	ErrNotStatusOk = errors.New("response status is not 200")
)
