package controllers

import (
	"goravel_by_gin/app/models"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
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
	var users []models.User
	if err := facades.Orm().Query().OrderByDesc("created_at").FindOrFail(&users); err != nil {
		facades.Log().Errorf("failed to get all users error message is: %s", err.Error())
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"message": "failed to find all users",
			"ok":      false,
		})
	}
	return ctx.Response().Success().Json(http.Json{
		"data": users,
		"ok":   true,
	})
}

// @Tags	User
// @Accept	json
// @Param	id	path	int	true	"USER ID"
// @Router	/users/{id} [GET]
func (r *UserController) Show(ctx http.Context) http.Response {
	return ctx.Response().Success().Json(http.Json{
		"Hello": "Show " + ctx.Request().Route("id"),
	})
}

// @Tags	User
// @Accept	json
// @Param	id		path	int						true	"USER ID"
// @Param	request	body	requests.UserRequest	true	"request body"
// @Router	/users/{id} [PUT]
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
// @Param	id	path	int	true	"USER ID"
// @Router	/users/{id} [DELETE]
func (r *UserController) Destroy(ctx http.Context) http.Response {
	return ctx.Response().Success().Json(http.Json{
		"message": "Destroy",
	})
}
