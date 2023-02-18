package gerrx

type ExHttpError interface {
	Error() string
	Response() (int, *CodeErrorResponse)
}
