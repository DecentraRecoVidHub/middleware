package response

import (
	"context"
	"encoding/json"
	"errors"
	Err "github.com/DecentraRecoVidHub/middleware/error"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func SuccessResponse(msg string, data interface{}) *Response {
	return &Response{200, msg, data}
}

func serverErrorResponse(err error) *Response {
	return &Response{http.StatusInternalServerError, err.Error(), nil}
}

func badRequestErrorResponse(err error) *Response {
	return &Response{http.StatusBadRequest, err.Error(), nil}
}

func unKnownErrorResponse(err error) *Response {
	return &Response{-1, "后端未正确包裹的错误：" + err.Error(), nil}
}

// 自定义的 ErrorEncoder 函数
func HttpErrorEncoder(ctx context.Context, err error, w http.ResponseWriter) {
	// 在错误响应中，Data 字段通常为空
	resp := unKnownErrorResponse(err)

	var e *Err.Err
	if errors.As(err, &e) {
		switch e.ErrorCode {
		case Err.BadRequest:
			resp = badRequestErrorResponse(err)
		case Err.ServerError:
			resp = serverErrorResponse(err)
		}

	}

	// 设置 HTTP 状态码和 Content-Type
	w.WriteHeader(resp.Code)
	w.Header().Set("Content-Type", "application/json")
	// 将 Response 结构体编码为 JSON 并写入响应体
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		// 处理 JSON 编码错误
		http.Error(w, "错误处理响应编码失败", http.StatusInternalServerError)
	}
}
