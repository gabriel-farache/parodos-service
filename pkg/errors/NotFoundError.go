package errors

import "fmt"

type NotFoundError struct {
	Message string
	Err     error
}

func (m NotFoundError) Error() string {
	if m.Err != nil {
		return fmt.Sprintf("NotFoundError Message: %q; err:%+v", m.Message, m.Err)
	}
	return m.Message
}
