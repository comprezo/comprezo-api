package handler

import (
	"comprezo/apperror"
	"encoding/json"
	"net/http"
)

type Handler struct {
	HandleFunc  HandleFunc
	ContentType string
}

type HandleFunc func(ctx Context) (interface{}, error)

type Context struct {
	Res http.ResponseWriter
	Req *http.Request
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func REST(handlerFunc HandleFunc) Handler {
	return Handler{
		HandleFunc:  handlerFunc,
		ContentType: "application/json",
	}
}

func (h Handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	ctx := Context{Res: res, Req: req}

	resp, err := h.HandleFunc(ctx)
	if err != nil {
		h.SendError(ctx, err)
		return
	}

	h.SendData(ctx, http.StatusOK, resp)
}

func (h Handler) SendData(ctx Context, code int, data interface{}) {
	ctx.Res.Header().Set("Content-Type", h.ContentType)
	ctx.Res.WriteHeader(code)

	json.NewEncoder(ctx.Res).Encode(data)
}

func (h Handler) SendError(ctx Context, err error) {
	code := http.StatusInternalServerError
	msg := "Internal Server Error"

	if appErr, ok := err.(apperror.AppError); ok {
		code = appErr.HTTPCode
		msg = appErr.PublicMsg
	}

	h.SendData(ctx, code, ErrorResponse{Error: msg})
}
