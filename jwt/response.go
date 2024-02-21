package jwt

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Response(ctx *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	ctx.JSON(httpStatus, gin.H{"code": code, "data": data, "msg": msg})
}

func Success(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusOK, 200, data, msg)
}

func Fail(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusBadRequest, 400, data, msg)
}

// 401
func FailUnauthorized(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusUnauthorized, 401, data, msg)
}

// 404
func FailStatusNotFound(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusNotFound, 404, data, msg)
}

// 409
func FailStatusConflict(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusConflict, 409, data, msg)
}

func NotAllowedMethod(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusMethodNotAllowed, 405, data, msg)
}

func CheckFail(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusUnprocessableEntity, 4022, data, msg)
}

func ServerError(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusInternalServerError, 500, data, msg)
}
