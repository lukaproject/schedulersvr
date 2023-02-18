package gerrx

import (
	"fmt"
	"net/http"
)

type ValidateError struct {
	SessionId   string
	Param2Error map[string]string
}

func (e *ValidateError) Error() string {
	msg := ""
	for param, e := range e.Param2Error {
		msg += fmt.Sprintf("param %s, error: %s\n", param, e)
	}
	return msg
}

func (e *ValidateError) Response() (int, *CodeErrorResponse) {
	return http.StatusBadRequest, &CodeErrorResponse{
		Msg:       e.Error(),
		SessionId: e.SessionId,
	}
}
