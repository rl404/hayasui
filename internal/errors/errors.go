package errors

import "errors"

// List of errors.
var (
	ErrRequiredToken    = errors.New("required discord token")
	ErrRequiredPrefix   = errors.New("required command prefix")
	ErrRequiredAPI      = errors.New("required api host")
	ErrRequiredRedis    = errors.New("required redis address")
	ErrInvalidCacheTime = errors.New("invalid caching time")
	ErrNot200           = errors.New("response code not 200")
)
