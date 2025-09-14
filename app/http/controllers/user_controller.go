package controllers

import (
	"github.com/goravel/framework/contracts/http"
)

type UserController struct {
	// Dependent services
}

func NewUserController() *UserController {
	return &UserController{
		// Inject services
	}
}

//	@Tags         User
//	@Accept       json
//	@Success      200
//	@Failure      400
//	@Router       /user [get]
func (r *UserController) Show(ctx http.Context) http.Response {
	return ctx.Response().Success().Json(http.Json{
		"Hello": "Goravel",
	})
}
