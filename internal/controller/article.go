package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"rest-api-service/internal/domain"
	"rest-api-service/internal/lib/utils"
	"rest-api-service/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type articleController struct {
	ctx            context.Context
	articleService service.ArticleServiceInterface
}

func InitArticleController(
	ctx context.Context,
	articleService service.ArticleServiceInterface,
	router *gin.RouterGroup,
) *articleController {
	articleController := &articleController{
		ctx:            ctx,
		articleService: articleService,
	}
	router.POST("/create/", articleController.CreateArticle)

	return articleController
}

// @Summary	Создание объявления
// @Tags		articles
// @Accept		json
// @Produce	json
// @Param		article	body		models.Article	true	"Объявление"
// @Success	201		{int}		id
// @Failure	400		{string}	string	"Barticle Request"
// @Failure	500		{string}	string	"Internal Server Error"
// @Router		/article/create/ [post]
func (a *articleController) CreateArticle(c *gin.Context) {
	article := &domain.Article{}
	err := json.NewDecoder(c.Request.Body).Decode(&article)
	if err != nil {
		utils.ErrorLog("unable to decode article", err)
		errorResponse(c, http.StatusBadRequest, err)
		return
	}

	validate := validator.New()
	err = validate.Struct(article)
	if err != nil {
		utils.ErrorLog("validate article", err)
		errorResponse(c, http.StatusBadRequest, err)
		return
	}

	id, err := a.articleService.CreateArticle(article)
	if err != nil {
		// if errors.Is(err, domain.ErrArticleNotFound) {
		// 	utils.ErrorLog("a.articleService.CreateArticle", err)
		// 	errorResponse(c, http.StatusNotFound, err)
		// 	return
		// }
		utils.ErrorLog("article creating by service", err)
		errorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}
