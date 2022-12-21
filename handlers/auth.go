package handlers

import (
	"UacademyGo/Blogpost/api_gateway/models"
	"UacademyGo/Blogpost/api_gateway/protogen/blogpost"
	"net/http"

	"github.com/gin-gonic/gin"
)

// //* AuthMyCORSMiddleware ...
func (h handler) AuthMiddleware() gin.HandlerFunc {
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
		ctx.Next()
	}
}