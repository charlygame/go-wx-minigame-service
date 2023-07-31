package utils

import "fmt"

type GameError struct {
	StatusCode int
	Err        error
}

func (e *GameError) Error() string {
	return fmt.Sprintf("status: %d, error: %s", e.StatusCode, e.Err.Error())
}
