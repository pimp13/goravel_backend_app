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

// @Tags	User
// @Accept	json
// @Router	/users [GET]
func (r *UserController) Index(ctx http.Context) http.Response {
	return ctx.Response().Success().Json(http.Json{
		"message": "Index",
	})
}

// @Tags	User
// @Accept	json
// @Param	id	path	int	true	"USER ID"
// @Router	/users/{id} [GET]
func (r *UserController) Show(ctx http.Context) http.Response {
	return ctx.Response().Success().Json(http.Json{
		"Hello": "Show",
	})
}

// @Tags	User
// @Accept	json
// @Param	id		path	int						true	"USER ID"
// @Param	request	body	requests.UserRequest	true	"request body"
// @Router	/users [PUT]
func (r *UserController) Update(ctx http.Context) http.Response {
	return ctx.Response().Success().Json(http.Json{
		"message": "Update",
	})
}


// @Tags	User
// @Accept	json
// @Router	/users [POST]
// @Param	request	body	requests.UserRequest	true	"request body"
func (r *UserController) Store(ctx http.Context) http.Response {
	return ctx.Response().Success().Json(http.Json{
		"Hello": "Store",
	})
}

// @Tags	User
// @Accept	json
// @Param	id		path	int						true	"USER ID"
// @Router	/users [DELETE]
func (r *UserController) Destroy(ctx http.Context) http.Response {
	return ctx.Response().Success().Json(http.Json{
		"message": "Destroy",
	})
}
