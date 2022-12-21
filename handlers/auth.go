package handlers

import (
	"UacademyGo/Blogpost/api_gateway/models"
	"UacademyGo/Blogpost/api_gateway/protogen/blogpost"
	"net/http"

	"github.com/gin-gonic/gin"
)

// //* AuthMyCORSMiddleware ...
func (h handler) AuthMiddleware(userType string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		hasAccesResponse, err := h.grpcClients.Auth.HasAcces(ctx.Request.Context(), &blogpost.TokenRequest{
			Token: token,
		})

		if err != nil {
			ctx.JSON(http.StatusBadRequest, models.JSONErrorRespons{
				Error: err.Error(),
			})
			ctx.Abort()
			return
		}

		if !hasAccesResponse.HasAcces {
			ctx.JSON(http.StatusUnauthorized, "Unauthorized")
			ctx.Abort()
			return
		}


		if userType!="*" {
			if hasAccesResponse.User.UserType != userType {
				ctx.JSON(http.StatusUnauthorized, "Permission Denied")
				ctx.Abort()
			}
		}

		ctx.Set("auth_username", hasAccesResponse.User.Username)
		ctx.Set("auth_username", hasAccesResponse.User.Id)
		
		ctx.Next()
	}
}


// * ================== Login ======================
// Login godoc
// @Summary     Login
// @Description Login
// @Tags        authorization (login)
// @Accept      json
// @Produce     json
// @Param       login body     models.LoginModel true "Login body"
// @Success     201   {object} models.JSONRespons{data=models.TokenResponse}
// @Failure     400   {object} models.JSONErrorRespons
// @Router      /v1/login [post]
func (h *handler) Login(ctx *gin.Context) {
	var body models.LoginModel
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorRespons{Error: err.Error()})
		return
	}

	tokenResponse, err := h.grpcClients.Auth.Login(ctx.Request.Context(), &blogpost.LoginRequest{
		Username: body.Username,
		Password: body.Password,
	}) 
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorRespons{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, models.JSONRespons{
		Message: "Article successfully created",
		Data:    tokenResponse,
	})
}