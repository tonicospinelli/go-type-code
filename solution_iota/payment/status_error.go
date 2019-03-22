package payment

import "fmt"

type errStatusNotAvailable struct {
	status interface{}
}

func (e errStatusNotAvailable) Error() string {
	return fmt.Sprintf("status %s is not avaible", e.status)
}
