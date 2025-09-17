package controllers

import (
	"goravel_by_gin/app/http/requests"
	"goravel_by_gin/app/services"
	"strconv"
	"time"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

type AuthController struct {
	// Dependent services
	authService services.AuthService
}

func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{
		// Inject services
		authService,
	}
}

// @Tags	Auth
// @Accept	json
// @Router	/auth [POST]
// @Param	request	body	requests.AuthRequest	true	"request body"
func (r *AuthController) Store(ctx http.Context) http.Response {
	var bodyData requests.AuthRequest

	validateErr, err := ctx.Request().ValidateRequest(&bodyData)
	if err != nil || validateErr != nil {
		return ctx.Response().Json(http.StatusBadRequest, http.Json{
			"message": "validation body request is failed!",
			"ok":      false,
			"errors":  validateErr.All(),
		})
	}

	user, err := r.authService.Register(ctx.Context(), &bodyData)
	if err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"message": err.Error(),
			"ok":      false,
		})
	}

	token, err := facades.Auth(ctx).Login(user)
	if err != nil {
		facades.Log().Errorf("failed to loggin user: %v", err)
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"message": "failed to loggin user",
			"ok":      true,
		})
	}

	return ctx.Response().Success().Json(http.Json{
		"ok":      true,
		"message": "reqister user is successfully!",
		"data":    token,
	})
}

// @Tags		Auth
// @Accept		json
// @Router		/auth/info [GET]
// @Security	ApiKeyAuth
func (r *AuthController) Show(ctx http.Context) http.Response {
	idStr, err := facades.Auth(ctx).ID()
	if err != nil {
		return ctx.Response().Json(http.StatusUnauthorized, http.Json{
			"ok":      false,
			"message": "failed to get current user!",
		})
	}
	id, _ := strconv.Atoi(idStr)

	user, err := r.authService.FindUserById(ctx.Context(), uint(id))
	if err != nil {
		return ctx.Response().Json(http.StatusUnauthorized, http.Json{
			"message": "failed to get current user",
			"ok":      false,
		})
	}

	return ctx.Response().Success().Json(http.Json{
		"ok":   true,
		"data": user,
	})
}

// @Tags	Auth
// @Accept	json
// @Router	/auth/login [POST]
// @Param	request	body	requests.AuthLoginRequest	true	"request body"
func (r *AuthController) Login(ctx http.Context) http.Response {
	var bodyData requests.AuthLoginRequest
	validateErr, err := ctx.Request().ValidateRequest(&bodyData)
	if err != nil || validateErr != nil {
		return ctx.Response().Json(http.StatusBadRequest, http.Json{
			"message": "validation body request is failed!",
			"ok":      false,
			"errors":  validateErr.All(),
		})
	}

	user, err := r.authService.Login(ctx.Context(), &bodyData)
	if err != nil {
		return ctx.Response().Json(http.StatusBadRequest, http.Json{
			"message": err.Error(),
			"ok":      false,
		})
	}

	token, err := facades.Auth(ctx).Login(user)
	if err != nil {
		facades.Log().Errorf("failed to loggin user: %v", err)
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"message": "failed to loggin user",
			"ok":      true,
		})
	}

	return ctx.Response().Cookie(http.Cookie{
		Name:     facades.Config().GetString("jwt.token_cookie"),
		Value:    token,
		Path:     "/",
		Domain:   facades.Config().GetString("http.url"),
		Secure:   true,
		HttpOnly: true,
	}).Json(http.StatusOK, http.Json{
		"message": "your are logged!",
		"ok":      true,
	})

}

// @Tags	Auth
// @Accept	json
// @Router	/auth/logout [POST]
func (r *AuthController) Logout(ctx http.Context) http.Response {
	// if err := facades.Auth(ctx).Logout(); err != nil {
	// 	facades.Log().Errorf("failed to user logout error message: %v", err)
	// 	return ctx.Response().Json(http.StatusInternalServerError, http.Json{
	// 		"message": "logout is failed!",
	// 		"ok":      false,
	// 	})
	// }
	ctx.Response().Cookie(http.Cookie{
		Name:     facades.Config().GetString("jwt.token_cookie"),
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Now().Add(-time.Hour), // تنظیم زمان انقضا به گذشته
	})
	return ctx.Response().Success().Json(http.Json{
		"message": "your are logouted!",
		"ok":      true,
	})
}
