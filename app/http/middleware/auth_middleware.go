package middleware

import (
	"goravel_by_gin/app/models"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

func AuthMiddleware() http.Middleware {
	return func(ctx http.Context) {
		// authHeader := ctx.Request().Header("Authorization", "")
		// if authHeader == "" {
		// 	ctx.Response().Json(http.StatusUnauthorized, http.Json{
		// 		"message": "token for authorization is required!",
		// 		"ok":      false,
		// 	})
		// }

		tokenStr := ctx.Request().Cookie(facades.Config().GetString("jwt.token_cookie"))
		if tokenStr == "" {
			// ctx.Request().Abort(http.StatusUnauthorized)
			ctx.Response().Json(http.StatusUnauthorized, http.Json{
				"message": "token for authorization is required!",
				"ok":      false,
			})
			return
		}
		facades.Log().Infof("tokenStr ===> %+v\n", tokenStr)

		payload, err := facades.Auth(ctx).Parse(tokenStr)
		if err != nil {
			// ctx.Request().Abort(http.StatusUnauthorized)
			ctx.Response().Json(http.StatusUnauthorized, http.Json{
				"message": "token is invalid!",
				"ok":      false,
			})
			return
		}
		facades.Log().Infof("payload ===> %+v\n", payload)

		exists, err := facades.Orm().Query().Model(&models.User{}).Where("id", payload.Key).Exists()
		if err != nil {
			// ctx.Request().Abort(http.StatusUnauthorized)
			ctx.Response().Json(http.StatusUnauthorized, http.Json{
				"message": "token is invalid!!",
				"ok":      false,
			})
			return
		}
		if !exists {
			ctx.Request().Abort(http.StatusUnauthorized)
		}
		ctx.Request().Next()
	}
}
