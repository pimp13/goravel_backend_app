package routes

import (
	admincontroller "goravel_by_gin/app/http/controllers/v1/admin"
	"goravel_by_gin/app/http/middleware"

	"github.com/goravel/framework/contracts/route"
)

func Admin(router route.Router) {
	router.Prefix("/admin").
		Middleware(middleware.AdminMiddleware()).
		Group(func(router route.Router) {

			authController := admincontroller.NewAuthController()
			router.Get("/info", authController.Index)

		})
}
