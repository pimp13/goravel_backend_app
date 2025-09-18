package admin

import (
	"github.com/goravel/framework/contracts/http"
)

type AuthController struct {
	// Dependent services
}

func NewAuthController() *AuthController {
	return &AuthController{
		// Inject services
	}
}

// @Tags		[Admin]
// @Accept		json
// @Router		/v1/admin/info [GET]
// @Security	ApiKeyAuth
func (r *AuthController) Index(ctx http.Context) http.Response {
	return ctx.Response().Success().Json(http.Json{
		"message": "Hello, welcome to admin panel.",
		"ok":      true,
	})
}
