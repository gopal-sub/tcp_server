package server


import (
	"errors"
)

var ServerShutdown = errors.New("Server shutdown")