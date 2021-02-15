package response

import "fmt"

type Error struct {
	Type string `json:"type,omitempty"`
	Msg  string `json:"msg,omitempty"`
	Code int    `json:"code,omitempty"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s - %s - %d", e.Type, e.Msg, e.Code)
}
