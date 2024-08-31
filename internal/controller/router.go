package controller

import (
	"context"
	"rest-api-service/internal/service"

	_ "rest-api-service/api/swagger"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/http-swagger"
)

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a rest-api-boilerplate.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/

//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/
func NewRouter(ctx context.Context, service *service.Service) *gin.Engine {
	router := gin.Default()

	articleRouter := router.Group("/article/")
	InitArticleController(ctx, service.Article, articleRouter)

	router.GET("/swagger/*any", gin.WrapF(httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	)))

	return router
}
