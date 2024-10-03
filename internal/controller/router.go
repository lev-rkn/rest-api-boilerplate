package controller

import (
	"context"
	"rest-api-service/internal/service"

	_ "rest-api-service/api/swagger"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/http-swagger"
)

//	@title			rest-api-boilerplate
//	@version		1.0
//	@description	This is a rest-api-boilerplate.
//	@host			localhost:8080
//	@BasePath		/
func NewRouter(ctx context.Context, service *service.Service) *gin.Engine {
	router := gin.Default()


	articleRouter := router.Group("/article/")
	InitArticleController(ctx, service.Article, articleRouter)

	router.GET("/swagger/*any", gin.WrapF(httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	)))

	return router
}
