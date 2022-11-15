package response

import "errors"

var (
	ErrBadRequest = errors.New("bad request")
	ErrInvalid    = errors.New("invalid request")
)
