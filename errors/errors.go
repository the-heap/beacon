package errors

import "fmt"

// Error allows us to create constant errors
type Error string

func (e Error) Error() string {
	return string(e)
}

// Wrap returns an Error with a prepended msg
func Wrap(msg string, err error) Error {
	return Error(fmt.Sprintf("%v: %v", msg, err))
}
