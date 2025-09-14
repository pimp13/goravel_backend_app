package routes

import (
	"github.com/goravel/framework/facades"

	"goravel_by_gin/app/http/controllers"
)

func Api() {
	userController := controllers.NewUserController()
	facades.Route().Get("/users/{id}", userController.Show)
}
