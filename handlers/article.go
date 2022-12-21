package handlers

import (
	"net/http"
	"strconv"

	"UacademyGo/Blogpost/api_gateway/models"
	"UacademyGo/Blogpost/api_gateway/protogen/blogpost"

	"github.com/gin-gonic/gin"
)

// * ================== Create Article ======================
// CreateArticle godoc
// @Summary     Create article
// @Description Create a new article
// @Tags        articles
// @Accept      json
// @Produce     json
// @Param       article       body     models.CreateModelArticle true  "article body"
// @Param       Authorization header   string                    false "Authorization"
// @Success     201           {object} models.JSONRespons{data=models.Article}
// @Failure     400           {object} models.JSONErrorRespons
// @Router      /v1/article [post]
func (h *handler) CreateArticle(ctx *gin.Context) {
	var body models.CreateModelArticle
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorRespons{Error: err.Error()})
		return
	}

	article, err := h.grpcClients.Article.CreateArticle(ctx.Request.Context(), &blogpost.CreateArticleRequest{
		Content: &blogpost.Content{
			Title: body.Title,
			Body: body.Body,
		},
		AuthorId: body.AuthorID,
	}) 
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorRespons{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, models.JSONRespons{
		Message: "Article successfully created",
		Data:    article,
	})
}

// * ==================== GetArticleById ====================
// GetArticleById godoc
// @Summary     get article by id
// @Description get a new article
// @Tags        articles
// @Accept      json
// @Param       id            path   string true  "Article ID"
// @Param       Authorization header string false "Authorization"
// @Produce     json
// @Success     200 {object} models.JSONRespons{data=models.GetByIDArticleModel}
// @Failure     404 {object} models.JSONErrorRespons
// @Router      /v1/article/{id} [get]
func (h *handler) GetArticleById(ctx *gin.Context) {
	idStr := ctx.Param("id")

	//TODO UUID validation

	article, err := h.grpcClients.Article.GetArticleByID(ctx.Request.Context(), &blogpost.GetArticleByIDRequest{
		Id: idStr,
	}) 
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorRespons{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.JSONRespons{
		Message: "passed successfully",
		Data:    article,
	})
}

// * ==================== GetArticleList ====================
// GetArticleList godoc
// @Summary     List articles
// @Description get articles
// @Tags        articles
// @Accept      json
// @Produce     json
// @Param       offset        query    int    false "0"
// @Param       limit         query    int    false "10"
// @Param       search        query    string false "smth"
// @Param       Authorization header   string false "Authorization"
// @Success     200           {object} models.JSONRespons{data=[]models.Article}
// @Router      /v1/article [get]
func (h *handler) GetArticleList(ctx *gin.Context) {

	offsetStr := ctx.DefaultQuery("offset", h.cfg.DefaultOffset)
	limitStr := ctx.DefaultQuery("limit", h.cfg.DefaultLimit)
	searchStr := ctx.DefaultQuery("search", "")

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorRespons{
			Error: err.Error(),
		})
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorRespons{
			Error: err.Error(),
		})
		return
	}

	articleList, err := h.grpcClients.Article.GetArticleList(ctx.Request.Context(), &blogpost.GetArticleListRequest{
		Offset: int32(offset),
		Limit: int32(limit),
		Search: searchStr,
	}) 
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorRespons{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.JSONRespons{
		Message: "OK",
		Data:    articleList,
	})
}

// * ==================== SearchArticleByMyUsername ====================
// SearchArticleByMyUsername godoc
// @Summary     List articles
// @Description get articles
// @Tags        articles
// @Accept      json
// @Produce     json
// @Param       offset        query    int    false "0"
// @Param       limit         query    int    false "10"
// @Param       Authorization header   string false "Authorization"
// @Success     200           {object} models.JSONRespons{data=[]models.Article}
// @Router      /v1/my_articles [get]
func (h *handler) SearchArticleByMyUsername(ctx *gin.Context) {

	offsetStr := ctx.DefaultQuery("offset", h.cfg.DefaultOffset)
	limitStr := ctx.DefaultQuery("limit", h.cfg.DefaultLimit)

	usernameRaw, ok := ctx.Get("auth_username")

	username, ok := usernameRaw.(string)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, "Unauthorized")
		return
	}

	searchStr := username

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorRespons{
			Error: err.Error(),
		})
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorRespons{
			Error: err.Error(),
		})
		return
	}

	articleList, err := h.grpcClients.Article.GetArticleList(ctx.Request.Context(), &blogpost.GetArticleListRequest{
		Offset: int32(offset),
		Limit: int32(limit),
		Search: searchStr,
	}) 
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorRespons{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.JSONRespons{
		Message: "OK",
		Data:    articleList,
	})
}

// * ==================== UpdateArticle ====================
// UpdateArticle godoc
// @Summary     Update article
// @Description Update a new article
// @Tags        articles
// @Accept      json
// @Param       article       body   models.UpdateArticleModel true  "updating article"
// @Param       Authorization header string                    false "Authorization"
// @Produce     json
// @Success     200 {object} models.JSONRespons{data=models.Article}
// @Failure     400 {object} models.JSONErrorRespons
// @Router      /v1/article [put]
func (h *handler) UpdateArticle(ctx *gin.Context) {
	var body models.UpdateArticleModel
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorRespons{Error: err.Error()})
		return
	}

	article, err := h.grpcClients.Article.UpdateArticle(ctx.Request.Context(), &blogpost.UpdateArticleRequest{
		Content: &blogpost.Content{
			Title: body.Title,
			Body: body.Body,
		},
		Id: body.ID,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorRespons{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.JSONRespons{
		Message: "Article successfully updated",
		Data:    article,
	})
}

// * ==================== DeleteArticle ====================
// DeleteArticle godoc
// @Summary     Delete article
// @Description delete article
// @Tags        articles
// @Accept      json
// @Param       id            path   string true  "Article ID"
// @Param       Authorization header string false "Authorization"
// @Produce     json
// @Success     200 {object} models.JSONRespons{data=models.Article}
// @Failure     400 {object} models.JSONErrorRespons
// @Router      /v1/article/{id} [delete]
func (h *handler) DeleteArticle(ctx *gin.Context) {
	idStr := ctx.Param("id")

	article, err := h.grpcClients.Article.DeleteArticle(ctx.Request.Context(), &blogpost.DeleteArticleRequest{
		Id: idStr,
	}) 
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorRespons{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusNotFound, models.JSONRespons{
		Message: "Article suucessfully deleted",
		Data:    article,
	})
}

// * ==================== PingPong ====================
func (h *handler) Pong(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
