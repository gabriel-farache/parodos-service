package errors

import "fmt"

type BadRequestError struct {
	Message string
	Err     error
}

func (m BadRequestError) Error() string {
	if m.Err != nil {
		return fmt.Sprintf("BadRequestError Message: %q; err:%+v", m.Message, m.Err)
	}
	return m.Message
}
