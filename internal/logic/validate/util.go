package validate

import "github.com/lukaproject/schedulersvr/gerrx"

func NewValidateError(sessionId string) *gerrx.ValidateError {
	return &gerrx.ValidateError{
		SessionId:   sessionId,
		Param2Error: make(map[string]string),
	}
}

func ReturnValidateError(err *gerrx.ValidateError) error {
	if len(err.Param2Error) == 0 {
		return nil
	}
	return err
}
