package main

import (
	"log"
	"net/http"

	"UacademyGo/Blogpost/api_gateway/clients"
	"UacademyGo/Blogpost/api_gateway/config"
	docs "UacademyGo/Blogpost/api_gateway/docs" // docs is generated by Swag CLI, you have to import it.
	"UacademyGo/Blogpost/api_gateway/handlers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// * @license.name  Apache 2.0
// * @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

func main() {

	cfg := config.Load()

	if cfg.Environment != "development" {
		gin.SetMode(gin.ReleaseMode)
	}

	docs.SwaggerInfo.Title = cfg.App
	docs.SwaggerInfo.Version = cfg.AppVersion

	r := gin.New()

	if cfg.Environment != "production" {
		r.Use(gin.Logger(), gin.Recovery())
	}

	//r.Use(MyCORSMiddleware())   //! * Agar opshi endpoitlarga qoyish kerak bolsa shunday ishlatiladi

	// r.GET("/ping", MyCORSMiddleware(), func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, gin.H{		//! Yoki specific endpointga ham qoshishimiz mumkun
	// 		"message": "pong",
	// 	})
	// })

	grpcClients, err := clients.NewGrpcClients(cfg)
	if err != nil {
		panic(err)
	}

	h := handlers.NewHandler(cfg, grpcClients)

	v1 := r.Group("/v1")
	{
		v1.Use(MyCORSMiddleware(), AuthMiddleware())
		v1.POST("/article", h.CreateArticle)
		v1.GET("/article/:id", h.GetArticleById)
		v1.GET("/article", h.GetArticleList)
		v1.PUT("/article", h.UpdateArticle)
		v1.DELETE("/article/:id", h.DeleteArticle)

		v1.GET("/ping", h.Pong) //*testing example the localhost

		v1.POST("author", h.CreateAuthor)
		v1.GET("/author/:id", h.GetAuthorById)
		v1.GET("/author", h.GetAuthorList)
		v1.PUT("/author", h.UpdateAuthor)
		v1.DELETE("/author/:id", h.DeleteAuthor)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(cfg.HTTPPort)
}

func MyCORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Println("MyCORSMiddleware...")
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Credentials", "true")
		ctx.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")
		ctx.Header("Access-Control-Allow-HEADERS", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-TOKEN, Authorization, accept, origin, Cache-Control, X-Requested-With")
		ctx.Header("Access-Control-Max-Age", "3600")

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(204)
			return
		}
		ctx.Next()
	}
}

// //* AuthMyCORSMiddleware ...
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")

		if token != "MyToken" {
			ctx.JSON(http.StatusUnauthorized, "Unauthorized")
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
