package controllers

import (
	"fmt"
	"goravel_by_gin/app/http/requests"
	"goravel_by_gin/app/models"
	"strconv"

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
	id := ctx.Request().Route("id")
	var category models.Category
	if err := facades.Orm().Query().Where("id", id).FirstOrFail(&category); err != nil {
		facades.Log().Errorf("failed to get category by #%v error message is: %v", id, err)

		return ctx.Response().Json(http.StatusNotFound, http.Json{
			"ok":      false,
			"message": fmt.Sprintf("failed to get category by #%v", id),
		})
	}

	return ctx.Response().Success().Json(http.Json{
		"ok":   true,
		"data": category,
	})
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
// @Param	id		path	uint							true	"Category ID"
// @Param	request	body	requests.UpdateCategoryRequest	true	"request body"
// @Router	/category/{id} [PUT]
func (r *CategoryController) Update(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	var bodyData requests.UpdateCategoryRequest
	validateErrs, err := ctx.Request().ValidateRequest(&bodyData)
	switch {
	case validateErrs != nil:
		return ctx.Response().Json(http.StatusBadRequest, http.Json{
			"message": "Validation body data is failed!",
			"ok":      false,
			"errors":  validateErrs.All(),
		})
	case err != nil:
		return ctx.Response().Json(http.StatusBadRequest, http.Json{
			"message": "Validation body data is failed!!",
			"ok":      false,
			"errors":  err,
		})
	}

	if _, err := facades.Orm().Query().Where("id", id).Update(&models.Category{
		Name:        bodyData.Name,
		IsActive:    bodyData.IsActive,
		Description: bodyData.Description,
		Slug:        bodyData.Slug,
	}); err != nil {
		facades.Log().Errorf("failed to update category by #%v => %v\n", id, err)
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"message": "failed to update category",
			"ok":      false,
		})
	}

	return ctx.Response().Success().Json(http.Json{
		"ok":      true,
		"message": "update category is successfully",
	})
}

// @Tags	Category
// @Accept	json
// @Param	id	path	int	true	"Category ID"
// @Router	/category/{id} [DELETE]
func (r *CategoryController) Destroy(ctx http.Context) http.Response {
	id, _ := strconv.Atoi(ctx.Request().Route("id"))

	if _, err := facades.Orm().Query().Where("id", id).Delete(&models.Category{}); err != nil {
		facades.Log().Errorf("failed to delete category by #%v => %v\n", id, err)
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"message": "failed to delete category",
			"ok":      false,
		})
	}

	return ctx.Response().Success().Json(http.Json{
		"message": "delete category is ok",
		"ok":      true,
	})
}

// @Tags	Category
// @Accept	json
// @Param	id		path	uint								true	"Category ID"
// @Param	request	body	requests.ChangeStatusCategoryDto	true	"request body"
// @Router	/category/{id}/change-status [PATCH]
func (r *CategoryController) ChnageStatus(ctx http.Context) http.Response {
	id, err := strconv.Atoi(ctx.Request().Route("id"))
	if err != nil {
		return ctx.Response().Json(http.StatusBadRequest, http.Json{
			"message": "failed to parse id parameter",
			"ok":      false,
		})
	}

	var bodyData requests.ChangeStatusCategoryDto
	validateErrs, err := ctx.Request().ValidateRequest(&bodyData)
	switch {
	case validateErrs != nil:
		return ctx.Response().Json(http.StatusBadRequest, http.Json{
			"message": "Validation body data is failed!",
			"ok":      false,
			"errors":  validateErrs.All(),
		})
	case err != nil:
		return ctx.Response().Json(http.StatusBadRequest, http.Json{
			"message": "Validation body data is failed!!",
			"ok":      false,
			"errors":  err,
		})
	}

	if _, err := facades.Orm().
		Query().
		Where("id", id).
		Model(&models.Category{}).
		Update("is_active", bodyData.IsActive); err != nil {
		facades.Log().Errorf("failed to change status category by #%v => %v\n", id, err)
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"message": "failed to change status category",
			"ok":      false,
		})
	}

	return ctx.Response().Success().Json(http.Json{
		"message": "change category status is successfully",
		"ok":      true,
	})
}
