package controllers

import (
	"goravel_by_gin/app/http/requests"
	"goravel_by_gin/app/models"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
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
	var categories []models.Category
	if err := facades.Orm().Query().OrderByDesc("created_at").FindOrFail(&categories); err != nil {
		facades.Log().Errorf("failed to get all categories error message is: %s", err.Error())

		return ctx.Response().Json(http.StatusNotFound, http.Json{
			"message": "failed to find all categories",
			"ok":      false,
		})
	}

	return ctx.Response().Success().Json(http.Json{
		"data": categories,
		"ok":   true,
	})
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
// @Router	/category [POST]
// @Param	request	body	requests.CategoryRequest	true	"request body"
func (r *CategoryController) Store(ctx http.Context) http.Response {
	var bodyData requests.CategoryRequest
	validateErrs, err := ctx.Request().ValidateRequest(&bodyData)
	if err != nil || validateErrs != nil {
		return ctx.Response().Json(http.StatusBadRequest, http.Json{
			"message": "failed to validation!",
			"ok":      false,
			"errors":  validateErrs.All(),
		})
	}

	if err := facades.Orm().Query().Create(&models.Category{
		Name:        bodyData.Name,
		IsActive:    bodyData.IsActive,
		Description: bodyData.Description,
		Slug:        bodyData.Slug,
	}); err != nil {
		facades.Log().Errorf("failed to create new category => %v\n", err)
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"message": "failed to create new category",
			"ok":      false,
		})
	}

	return ctx.Response().Json(http.StatusCreated, http.Json{
		"message": "create new category is sucessfully!",
		"ok":      true,
	})
}

// @Tags	Category
// @Accept	json
// @Param	id		path	uint						true	"Category ID"
// @Param	request	body	requests.CategoryRequest	true	"request body"
// @Router	/category/{id} [PUT]
func (r *CategoryController) Update(ctx http.Context) http.Response {
	//var bodyData requests.CategoryRequest
	return nil
}

// @Tags	Category
// @Accept	json
// @Param	id	path	int	true	"Category ID"
// @Router	/category/{id} [DELETE]
func (r *CategoryController) Destroy(ctx http.Context) http.Response {
	return nil
}
