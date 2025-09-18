package controllers

import (
	"goravel_by_gin/app/http/requests"
	"goravel_by_gin/app/services"
	"strconv"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

type PostController struct {
	// Dependent services
	postService services.PostService
}

func NewPostController(postService services.PostService) *PostController {
	return &PostController{
		// Inject services
		postService,
	}
}

// @Tags		Post
// @Accept		json
// @Router		/post [GET]
// @Security	ApiKeyAuth
func (r *PostController) Index(ctx http.Context) http.Response {
	posts, err := r.postService.FindAll(ctx.Context())
	if err != nil {
		return ctx.Response().Json(http.StatusNotFound, http.Json{
			"message": err,
			"ok":      false,
		})
	}

	return ctx.Response().Success().Json(http.Json{
		"ok":   true,
		"data": posts,
	})
}

// @Tags	Post
// @Accept	json
// @Param	id	path	uint	true	"Post ID"
// @Router	/post/{id} [GET]
func (r *PostController) Show(ctx http.Context) http.Response {
	id, err := strconv.Atoi(ctx.Request().Route("id"))
	if err != nil {
		return ctx.Response().Json(http.StatusBadRequest, http.Json{
			"message": "failed to parse id parameter",
			"ok":      false,
		})
	}
	post, err := r.postService.FindById(ctx.Context(), uint(id))
	if err != nil {
		return ctx.Response().Json(http.StatusNotFound, http.Json{
			"message": err,
			"ok":      false,
		})
	}

	return ctx.Response().Success().Json(http.Json{
		"data": post,
		"ok":   true,
	})
}

// @Tags	Post
// @Accept	json
// @Router	/post [POST]
// @Param	request	body	requests.PostRequest	true	"request body"
func (r *PostController) Store(ctx http.Context) http.Response {
	var bodyData requests.PostRequest
	vaidateErrs, err := ctx.Request().ValidateRequest(&bodyData)
	if err != nil || vaidateErrs != nil {
		return ctx.Response().Json(http.StatusBadRequest, http.Json{
			"message": "validation is failed!",
			"ok":      false,
			"errors":  vaidateErrs.All(),
		})
	}

	file, err := ctx.Request().File("image")
	if err != nil {
		return ctx.Response().Json(http.StatusBadRequest, http.Json{
			"message": "image file is required",
		})
	}

	key, err := facades.Auth(ctx).ID()
	if err != nil {
		return ctx.Response().Json(http.StatusBadRequest, http.Json{
			"message": "failed to get current user!",
			"ok":      false,
		})
	}
	userId, _ := strconv.Atoi(key)

	if err := r.postService.Create(ctx.Context(), bodyData, uint(userId), file); err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"message": err.Error(),
			"ok":      false,
		})
	}

	return ctx.Response().Json(http.StatusCreated, http.Json{
		"message": "create new post is sucessfully!",
		"ok":      true,
	})
}

// @Tags	Post
// @Accept	json
// @Param	id		path	uint					true	"Post ID"
// @Param	request	body	requests.PostRequest	true	"request body"
// @Router	/post/{id} [PUT]
func (r *PostController) Update(ctx http.Context) http.Response {
	return nil
}

// @Tags	Post
// @Accept	json
// @Param	id	path	int	true	"Post ID"
// @Router	/post/{id} [DELETE]
func (r *PostController) Destroy(ctx http.Context) http.Response {
	return nil
}
