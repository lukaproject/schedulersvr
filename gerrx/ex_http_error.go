package gerrx

type ExHttpError interface {
	Error() string
	Response() (int, *CodeErrorResponse)
}

type DefaultHttpError struct {
	StatusCode int
	Msg        string
	SessionId  string
}

func (dhe *DefaultHttpError) Error() string {
	return dhe.Msg
}

func (dhe *DefaultHttpError) Response() (int, *CodeErrorResponse) {
	return dhe.StatusCode, &CodeErrorResponse{
		SessionId: dhe.SessionId,
		Msg:       dhe.Msg,
	}
}

func NewDefaultHttpError(statusCode int, msg string) *DefaultHttpError {
	return &DefaultHttpError{
		StatusCode: statusCode,
		Msg:        msg,
	}
}
