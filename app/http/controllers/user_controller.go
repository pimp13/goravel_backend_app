package controllers

import (
	"goravel_by_gin/app/http/requests"
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
	if err := facades.Orm().
		Query().
		With("Posts").
		OrderByDesc("created_at").
		FindOrFail(&users); err != nil {
		facades.Log().Errorf("failed to get all users error message is: %s", err.Error())
		return ctx.Response().Json(http.StatusNotFound, http.Json{
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
	var user models.User
	// id := strconv.Atoi(ctx.Request().Route("id"))

	if err := facades.Orm().Query().FindOrFail(&user, ctx.Request().Route("id")); err != nil {
		return ctx.Response().Json(http.StatusNotFound, http.Json{
			"message": "User not found",
			"ok":      false,
		})
	}
	return ctx.Response().Success().Json(http.Json{
		"data": user,
		"ok":   true,
	})
}

// @Tags	User
// @Accept	json
// @Param	id		path	int						true	"USER ID"
// @Param	request	body	requests.UserRequest	true	"request body"
// @Router	/users/{id} [PUT]
func (r *UserController) Update(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")

	var request requests.UserRequest
	// validateErrs, err := ctx.Request().Validate(request.Rules(ctx))
	// if err != nil || validateErrs.Fails() {
	// 	return ctx.Response().Json(http.StatusBadRequest, http.Json{
	// 		"message": "Validation error",
	// 		"ok":      false,
	// 		"errors":  validateErrs.Errors().All(),
	// 	})
	// }

	if err := ctx.Request().Bind(&request); err != nil {
		facades.Log().Errorf("failed to binding request for update user by #%v: %v\n", id, err)

		return ctx.Response().Json(http.StatusBadRequest, http.Json{
			"message": err.Error(),
			"ok":      false,
		})
	}

	if _, err := facades.Orm().Query().Where("id", id).Update(&models.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}); err != nil {
		facades.Log().Errorf("failed to update user by #%v: %v\n", id, err)
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"message": "Update error",
			"ok":      false,
		})
	}

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
