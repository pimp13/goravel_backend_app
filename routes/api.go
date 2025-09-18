package routes

import (
	"log"

	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"

	"goravel_by_gin/app/http/controllers"
	"goravel_by_gin/app/http/middleware"
	"goravel_by_gin/app/services"
)

func Api() {

	facades.Route().Prefix("/api").Group(func(router route.Router) {

		//* users service
		userController := controllers.NewUserController()
		router.Resource("/users", userController)

		//* category service
		categoryController := controllers.NewCategoryController()
		router.Resource("/category", categoryController)
		router.Patch("/category/{id}/change-status", categoryController.ChnageStatus)

		//* post service
		postService, err := facades.App().Make("services.PostService")
		if err != nil {
			log.Fatalf("Failed to make post service: %v", err)
		}
		postController := controllers.NewPostController(postService.(services.PostService))
		router.Resource("/post", postController)

		//* auth service
		authService, _ := facades.App().Make("services.AuthService")
		authController := controllers.NewAuthController(authService.(services.AuthService))
		router.Post("/auth", authController.Store)
		router.Post("/auth/login", authController.Login)
		router.Post("/auth/logout", authController.Logout)
		router.Middleware(middleware.AuthMiddleware()).Get("/auth/info", authController.Show)

		//* V1 Routes
		router.Prefix("/v1").Group(func(router route.Router) {
			//* Admin Routes
			Admin(router)
		})

	})
}
