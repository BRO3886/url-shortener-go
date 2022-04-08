package shortener

import "errors"

var (
	ErrRedirectNotFound = errors.New("redirect not found")
	ErrRedirectInvalid  = errors.New("redirect invalid")
)
