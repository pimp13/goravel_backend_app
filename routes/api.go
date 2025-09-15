package routes

import (
	"github.com/goravel/framework/facades"

	"goravel_by_gin/app/http/controllers"
)

func Api() {

	facades.Route().Resource(
		"/users",
		controllers.NewUserController(),
	)

	facades.Route().Resource(
		"/category",
		controllers.NewCategoryController(),
	)

}
