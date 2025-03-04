package handlers

import (
	"comprezo/apperror"
	"comprezo/handler"
	"net/http"
	"net/url"
	"strconv"
)

type GetSizeResponse struct {
	Size int64 `json:"size"`
}

func Home(ctx handler.Context) (interface{}, error) {
	return map[string]string{"message": "OK!"}, nil
}

func GetSize(ctx handler.Context) (interface{}, error) {
	queryParams := ctx.Req.URL.Query()
	targetURL := queryParams.Get("url")

	if targetURL == "" {
		return nil, apperror.New(nil, http.StatusBadRequest, "Missing url parameter")
	}

	_, err := url.ParseRequestURI(targetURL)
	if err != nil {
		return nil, apperror.New(err, http.StatusBadRequest, "Invalid url")
	}

	resp, err := http.Head(targetURL)
	if err != nil {
		return nil, apperror.New(err, http.StatusInternalServerError, "Failed to fetch url")
	}
	defer resp.Body.Close()

	contentLengthStr := resp.Header.Get("Content-Length")
	contentLength, err := strconv.ParseInt(contentLengthStr, 10, 64)
	if err != nil {
		contentLength = 0
	}

	return GetSizeResponse{Size: contentLength}, nil
}
