package controllers

import (
	"github.com/goravel/framework/contracts/http"
)

type CategoryController struct {
	// Dependent services
}

func NewCategoryController() *CategoryController {
	return &CategoryController{
		// Inject services
	}
}

// @Tags	Category
// @Accept	json
// @Router	/category [GET]
func (r *CategoryController) Index(ctx http.Context) http.Response {
	return ctx.Response().Success().Json("hello world")
}

// @Tags	Category
// @Accept	json
// @Param	id	path	uint	true	"Category ID"
// @Router	/category/{id} [GET]
func (r *CategoryController) Show(ctx http.Context) http.Response {
	return nil
}

// @Tags	Category
// @Accept	json
// @Param	id		path	uint						true	"Category ID"
// @Param	request	body	requests.CategoryRequest	true	"request body"
// @Router	/category/{id} [PUT]
func (r *CategoryController) Store(ctx http.Context) http.Response {
	return nil
}

// @Tags	Category
// @Accept	json
// @Router	/category [POST]
// @Param	request	body	requests.CategoryRequest	true	"request body"
func (r *CategoryController) Update(ctx http.Context) http.Response {
	return nil
}

// @Tags	Category
// @Accept	json
// @Param	id	path	int	true	"Category ID"
// @Router	/category/{id} [DELETE]
func (r *CategoryController) Destroy(ctx http.Context) http.Response {
	return nil
}
