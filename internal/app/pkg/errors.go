package pkg

import (
	"errors"
	"fmt"
)

var ErrUnknownCy = fmt.Errorf("unknown currency")

var (
	ErrNoUser            = errors.New("user is not found")
	ErrInsufficientFunds = errors.New("insufficient funds")
)

var ErrInvalidInput = errors.New("input validation failed")


