package gerrx

import (
	"context"
	"net/http"
)

type CodeError struct {
	Code      int    `json:"code"`
	Msg       string `json:"msg"`
	SessionId string `json:"session_id,optional"`
}

type CodeErrorResponse struct {
	Code      int    `json:"code"`
	Msg       string `json:"msg"`
	SessionId string `json:"session_id,optional"`
}

func NewCodeError(code int, msg string, session_id string) error {
	return &CodeError{Code: code, Msg: msg, SessionId: session_id}
}

func NewDefaultError(msg string, session_id string) error {
	return NewCodeError(0, msg, session_id)
}

func (e *CodeError) Error() string {
	return e.Msg
}

func (e *CodeError) Data() *CodeErrorResponse {
	return &CodeErrorResponse{
		Code:      e.Code,
		Msg:       e.Msg,
		SessionId: e.SessionId,
	}
}

func ErrorHandlerCtx(_ context.Context, err error) (int, interface{}) {
	switch e := err.(type) {
	case *CodeError:
		return http.StatusOK, e.Data()
	default:
		return http.StatusInternalServerError, nil
	}
}
