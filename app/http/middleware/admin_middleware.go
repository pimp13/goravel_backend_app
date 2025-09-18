package middleware

import (
	"goravel_by_gin/app/models"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

func AdminMiddleware() http.Middleware {
	return func(ctx http.Context) {
		tokenStr := ctx.Request().Cookie(facades.Config().GetString("jwt.token_cookie"))
		if tokenStr == "" {
			ctx.Request().AbortWithStatusJson(http.StatusUnauthorized, http.Json{
				"message": "you is not login!",
				"ok":      false,
			})
			return
		}

		payload, err := facades.Auth(ctx).Parse(tokenStr)
		if err != nil {
			ctx.Request().AbortWithStatusJson(http.StatusUnauthorized, http.Json{
				"message": "token is invalid!",
				"ok":      false,
			})
			return
		}

		var user models.User
		if err := facades.Orm().
			Query().
			Model(&models.User{}).
			Where("id", payload.Key).
			FirstOrFail(&user); err != nil {
			ctx.Request().AbortWithStatusJson(http.StatusUnauthorized, http.Json{
				"message": "token is invalid!",
				"ok":      false,
			})
			return
		}

		if user.ID == 0 {
			ctx.Request().AbortWithStatusJson(http.StatusUnauthorized, http.Json{
				"message": "token is invalid!",
				"ok":      false,
			})
			return
		}
		if user.IsSupperAdmin && user.IsActive {
			ctx.Request().Next()
			return
		}
		if !(user.Role.IsAdmin() && user.IsActive) {
			ctx.Request().AbortWithStatusJson(http.StatusForbidden, http.Json{
				"message": "you can't permission!",
				"ok":      false,
			})
			return
		}

		ctx.Request().Next()
	}
}
