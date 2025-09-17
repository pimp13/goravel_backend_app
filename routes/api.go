package routes

import (
	"log"

	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"

	"goravel_by_gin/app/http/controllers"
	"goravel_by_gin/app/services"
)

func Api() {

	facades.Route().Prefix("/api").Group(func(router route.Router) {

		userController := controllers.NewUserController()
		router.Resource("/users", userController)

		categoryController := controllers.NewCategoryController()
		router.Resource("/category", categoryController)
		router.Patch("/category/{id}/change-status", categoryController.ChnageStatus)

		// uploadService, err := facades.App().Make("services.UploadService")
		// if err != nil {
		// 	log.Fatalf("Failed to make upload service: %v", err)
		// }
		postService, err := facades.App().Make("services.PostService")
		if err != nil {
			log.Fatalf("Failed to make post service: %v", err)
		}
		postController := controllers.NewPostController(postService.(services.PostService))
		router.Resource("/post", postController)

	})
}
