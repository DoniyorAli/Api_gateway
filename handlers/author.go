package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"UacademyGo/Blogpost/api_gateway/models"
	"UacademyGo/Blogpost/api_gateway/protogen/blogpost"

	"github.com/gin-gonic/gin"
)

// * ================== Create Author =========================
// CreateAuthor godoc
// @Summary     Create author
// @Description Create a new author
// @Tags        authors
// @Accept      json
// @Param       author        body   models.CreateModelAuthor true  "author body"
// @Param       Authorization header string                   false "Authorization"
// @Produce     json
// @Success     201 {object} models.JSONRespons{data=string}
// @Failure     400 {object} models.JSONErrorRespons
// @Router      /v1/author [post]
func (h *handler) CreateAuthor(ctx *gin.Context) {
	var body models.CreateModelAuthor
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorRespons{Error: err.Error()})
		return
	}

	obj, err := h.grpcClients.Author.CreateAuthor(ctx.Request.Context(), &blogpost.CreateAuthorRequest{
		Fullname: body.Fullname,
		Middlename: body.Middlename,
	}) 
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorRespons{
			Error: err.Error(),
		})
		return
	}

	article, err := h.grpcClients.Author.GetAuthorByID(ctx.Request.Context(), &blogpost.GetAuthorByIDRequest{
		Id: obj.Id,
	}) 
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorRespons{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, models.JSONRespons{
		Message: "Author | GetList",
		Data:    article,
	})
}

// * ==================== Get Author By Id ====================
// GetAuthorById godoc
// @Summary     get author by id
// @Description get a new author
// @Tags        authors
// @Accept      json
// @Param       id            path   string true  "Article ID"
// @Param       Authorization header string false "Authorization"
// @Produce     json
// @Success     200 {object} models.JSONRespons{data=models.Author}
// @Failure     404 {object} models.JSONErrorRespons
// @Router      /v1/author/{id} [get]
func (h *handler) GetAuthorById(ctx *gin.Context) {
	idStr := ctx.Param("id")

	//TODO UUID validation

	author, err := h.grpcClients.Author.GetAuthorByID(ctx.Request.Context(), &blogpost.GetAuthorByIDRequest{
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
		Data:    author,
	})
}

// * ==================== Get Author List ====================
// GetArticleList godoc
// @Summary     List authors
// @Description get authors
// @Tags        authors
// @Accept      json
// @Produce     json
// @Param       offset        query    int    false "0"
// @Param       limit         query    int    false "10"
// @Param       search        query    string false "smth"
// @Param       Authorization header   string false "Authorization"
// @Success     200           {object} models.JSONRespons{data=[]models.Author}
// @Router      /v1/author [get]
func (h *handler) GetAuthorList(ctx *gin.Context) {
	offsetStr := ctx.DefaultQuery("offset", "0")
	limitStr := ctx.DefaultQuery("limit", "10")
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

	authorList, err := h.grpcClients.Author.GetAuthorList(ctx.Request.Context(), &blogpost.GetAuthorListRequest{
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
		Data:    authorList,
	})
}

// * ==================== Update Author =======================
// UpdateAuthor godoc
// @Summary     Update author
// @Description Update a new author
// @Tags        authors
// @Accept      json
// @Param       author        body   models.UpdateAuthorResponse true  "updating author"
// @Param       Authorization header string                      false "Authorization"
// @Produce     json
// @Success     200 {object} models.JSONRespons{data=models.Author}
// @Failure     400 {object} models.JSONErrorRespons
// @Router      /v1/author [put]
func (h *handler) UpdateAuthor(ctx *gin.Context) {
	var body models.UpdateAuthorResponse
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorRespons{Error: err.Error()})
		return
	}

	_, err := h.grpcClients.Author.UpdateAuthor(ctx.Request.Context(), &blogpost.UpdateAuthorRequest{
		Id: body.ID,
		Fullname: body.Fullname,
		Middlename: body.Middlename,
	})
	fmt.Println("ssss",err)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorRespons{
			Error: err.Error(),
		})
		return
	}
	
	author, err := h.grpcClients.Author.GetAuthorByID(ctx.Request.Context(), &blogpost.GetAuthorByIDRequest{
		Id: body.ID,
	}) 
	fmt.Println("ssssxxx",err)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorRespons{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.JSONRespons{
		Message: "Author successfully updated",
		Data:    author,
	})
}

// * ==================== Delete Author =======================
// DeleteAuthor godoc
// @Summary     Delete author
// @Description delete author
// @Tags        authors
// @Accept      json
// @Param       id            path   string true  "Author ID"
// @Param       Authorization header string false "Authorization"
// @Produce     json
// @Success     200 {object} models.JSONRespons{data=models.Author}
// @Failure     400 {object} models.JSONErrorRespons
// @Router      /v1/author/{id} [delete]
func (h *handler) DeleteAuthor(ctx *gin.Context) {
	idStr := ctx.Param("id")

	obj, err := h.grpcClients.Author.GetAuthorByID(ctx.Request.Context(), &blogpost.GetAuthorByIDRequest{
		Id: idStr,
	}) 
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorRespons{
			Error: err.Error(),
		})
		return
	}

	author, err := h.grpcClients.Author.DeleteAuthor(ctx.Request.Context(), &blogpost.DeleteAuthorRequest{
		Id: obj.Id,
	}) 
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorRespons{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusNotFound, models.JSONRespons{
		Message: "Author suucessfully deleted",
		Data:    author,
	})
}
