package res

import (
	"github.com/goravel/framework/contracts/http"
)

func ErrorResponse(ctx http.Context, message string, status ...int) http.Response {
	code := http.StatusInternalServerError
	if status[0] != 0 {
		code = status[0]
	}

	return ctx.Response().Json(code, http.Json{
		"message": message,
		"ok":      false,
	})
}

func SuccessResponse(ctx http.Context, data any, status ...int) http.Response {
	code := http.StatusOK
	if status[0] != 0 {
		code = status[0]
	}

	return ctx.Response().Json(code, http.Json{
		"ok":   true,
		"data": data,
	})
}
