package protocol

import (
	"errors"
)

var InvalidCommand = errors.New("Invalid command")
var NoCommandProvided = errors.New("No command")