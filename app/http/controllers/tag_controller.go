package controllers

import (
	"goravel_by_gin/app/services"

	"github.com/goravel/framework/contracts/http"
)

type TagController struct {
	// Dependent services
	tagService services.TagService
}

func NewTagController(tagService services.TagService) *TagController {
	return &TagController{
		// Inject services
		tagService,
	}
}

// @Tags		[Tag]
// @Accept		json
// @Router		/tag [GET]
// @Security	ApiKeyAuth
func (r *TagController) Index(ctx http.Context) http.Response {
	tags, err := r.tagService.FindAll(ctx.Context())
	if err != nil {
		return ctx.Response().Json(http.StatusNotFound, http.Json{
			"message": err.Error(),
			"ok":      false,
		})
	}
	return ctx.Response().Json(http.StatusNotFound, http.Json{
		"data": tags,
		"ok":   true,
	})
}

// @Tags	[Tag]
// @Accept	json
// @Param	id	path	uint	true	"Tag ID"
// @Router	/tag/{id} [GET]
// @Security	ApiKeyAuth
func (r *TagController) Show(ctx http.Context) http.Response {
	return nil
}

// @Tags	[Tag]
// @Accept	json
// @Router	/tag [POST]
// @Param	request	body	requests.TagRequest	true	"request body"
// @Security	ApiKeyAuth
func (r *TagController) Store(ctx http.Context) http.Response {
	return nil
}

// @Tags	[Tag]
// @Accept	json
// @Param	id		path	uint					true	"Tag ID"
// @Param	request	body	requests.TagRequest	true	"request body"
// @Router	/tag/{id} [PUT]
// @Security	ApiKeyAuth
func (r *TagController) Update(ctx http.Context) http.Response {
	return nil
}

// @Tags	[Tag]
// @Accept	json
// @Param	id	path	int	true	"Tag ID"
// @Router	/tag/{id} [DELETE]
// @Security	ApiKeyAuth
func (r *TagController) Destroy(ctx http.Context) http.Response {
	return nil
}
