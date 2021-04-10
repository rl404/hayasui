package internal

import "errors"

var (
	errRequiredToken    = errors.New("required discord token")
	errRequiredPrefix   = errors.New("required command prefix")
	errRequiredAPI      = errors.New("required api host")
	errRequiredRedis    = errors.New("required redis address")
	errInvalidCacheTime = errors.New("invalid caching time")
	errNot200           = errors.New("response code not 200")
)
