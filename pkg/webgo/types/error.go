package types

import "fmt"

type Error struct {
	StatusCode int    `json:"status"`
	Reason     string `json:"reason"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("%d, %s", e.StatusCode, e.Reason)
}
