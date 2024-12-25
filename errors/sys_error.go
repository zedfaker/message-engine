package syserror

import "github.com/pkg/errors"

type InitError struct {
	Err error
}

func (e *InitError) Error() string {
	return e.Err.Error()
}

func NewInitError(errorMessage string, err error) *InitError {
	return &InitError{
		Err: errors.Wrap(err, errorMessage),
	}
}
